// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PostStorage {
    enum Category {
        Blog,
        News,
        Article
    }

    struct Post {
        address author;
        string title;
        string content;
        uint256 timestamp;
        Category category;
    }

    mapping(address => Post[]) public userPosts;
    address[] public users;

    event PostPublished(
        address indexed author,
        string title,
        Category category
    );

    function post(
        string memory _title,
        string memory _content,
        Category _category,
        address author
    ) public {
        if (userPosts[author].length == 0) {
            users.push(author);
        }

        Post memory newPost = Post({
            author: author,
            content: _content,
            title: _title,
            timestamp: block.timestamp,
            category: _category
        });

        userPosts[author].push(newPost);

        emit PostPublished(author, _title, _category);
    }

    function getUsersPost(address user) public view returns (Post[] memory) {
        return userPosts[user];
    }

    function getPostByIndex(
        address user,
        uint256 index
    ) public view returns (Post memory) {
        require(index < userPosts[user].length, "Post does not exist");
        return userPosts[user][index];
    }

    function getAllPosts() public view returns (Post[] memory) {
        uint totalPostsCount = 0;
        for (uint i = 0; i < users.length; i++) {
            totalPostsCount += userPosts[users[i]].length;
        }

        Post[] memory allPosts = new Post[](totalPostsCount);
        uint currentIndex = 0;
        for (uint i = 0; i < users.length; i++) {
            Post[] memory userPostArray = userPosts[users[i]];
            for (uint j = 0; j < userPostArray.length; j++) {
                allPosts[currentIndex] = userPostArray[j];
                currentIndex++;
            }
        }
        return allPosts;
    }

    function getAllUsers() public view returns (address[] memory) {
        return users;
    }
}
