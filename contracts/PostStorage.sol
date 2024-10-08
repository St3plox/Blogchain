// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PostStorage {
    enum Category {
        Blog,
        News,
        Article
    }

    struct Post {
        uint256 id;
        address author;
        string title;
        string content;
        string[] tags;
        uint256 timestamp;
        Category category;
        string[] media_names;
        string[] media_urls;
    }

    mapping(address => Post[]) public userPosts;
    mapping(uint256 => Post) private idToPost;
    address[] public users;
    uint256 public postCounter;

    event PostPublished(
        uint256 id,
        address indexed author,
        string title,
        string[] tags,
        Category category
    );

    function post(
        string memory _title,
        string memory _content,
        Category _category,
        string[] memory _tags,
        address author,
        string[] memory _media_names,
        string[] memory _media_urls
    ) public {
        require(_media_names.length == _media_urls.length, "Media names and URLs array length mismatch");

        if (userPosts[author].length == 0) {
            users.push(author);
        }

        postCounter++;
        uint256 newId = postCounter;

        Post memory newPost = Post({
            id: newId,
            author: author,
            content: _content,
            title: _title,
            timestamp: block.timestamp,
            tags: _tags,
            category: _category,
            media_names: _media_names,
            media_urls: _media_urls
        });

        userPosts[author].push(newPost);
        idToPost[newId] = newPost;

        emit PostPublished(newId, author, _title, _tags, _category);
    }

    function getPostByID(uint256 id) public view returns (Post memory) {
        require(idToPost[id].id != 0, "Post with this ID does not exist");
        return idToPost[id];
    }

    function getUsersPost(address user) public view returns (Post[] memory) {
        return userPosts[user];
    }

    function getUsersPostPaginated(
        address user,
        uint256 page,
        uint256 pageSize
    ) public view returns (Post[] memory) {
        uint256 totalPosts = userPosts[user].length;
        uint256 start = page * pageSize;
        uint256 end = start + pageSize;

        if (end > totalPosts) {
            end = totalPosts;
        }

        require(start < totalPosts, "Page out of range");

        Post[] memory paginatedPosts = new Post[](end - start);
        for (uint256 i = start; i < end; i++) {
            paginatedPosts[i - start] = userPosts[user][i];
        }

        return paginatedPosts;
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

    function getPostsPaginated(
        uint256 page,
        uint256 pageSize
    ) public view returns (Post[] memory) {
        uint totalPostsCount = 0;
        for (uint i = 0; i < users.length; i++) {
            totalPostsCount += userPosts[users[i]].length;
        }

        require(page * pageSize < totalPostsCount, "Page out of range");

        uint start = page * pageSize;
        uint end = start + pageSize > totalPostsCount
            ? totalPostsCount
            : start + pageSize;

        Post[] memory paginatedPosts = new Post[](end - start);
        uint currentIndex = 0;
        uint postIndex = 0;
        for (uint i = 0; i < users.length; i++) {
            Post[] memory userPostArray = userPosts[users[i]];
            for (uint j = 0; j < userPostArray.length; j++) {
                if (postIndex >= start && postIndex < end) {
                    paginatedPosts[currentIndex] = userPostArray[j];
                    currentIndex++;
                }
                postIndex++;
            }
        }

        return paginatedPosts;
    }

    function getAllUsers() public view returns (address[] memory) {
        return users;
    }
}
