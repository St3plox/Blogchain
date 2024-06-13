// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./PostCategory.sol";

contract PostStorage {
    struct Post {
        address author;
        string title;
        string content;
        uint256 timestamp;
        Category category;
    }

    mapping(address => Post[]) public userPosts;

    event PostPublished(
        address indexed author,
        string title,
        Category category
    );

    function post(string memory _title, string memory _content, Category _category) public {
        
        Post memory newPost = Post({
            author: msg.sender,
            content: _content,
            title: _title,
            timestamp: block.timestamp,
            category: _category
        });

        userPosts[msg.sender].push(newPost);

        emit PostPublished(msg.sender, _title, _category);
    }

    function getMyPosts() public view returns (Post[] memory) {
        return userPosts[msg.sender];
    }

    function getUsersPost(address user)public view returns (Post[] memory){
        return userPosts[user];
    }
    
}
