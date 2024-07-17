// test/PostStorageTest.js
const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("PostStorage", function () {
    let PostStorage;
    let postStorage;
    let owner;
    let addr1;
    let addr2;

    beforeEach(async function () {
        [owner, addr1, addr2] = await ethers.getSigners();
        const PostStorageFactory = await ethers.getContractFactory("PostStorage");
        postStorage = await PostStorageFactory.deploy();
        await postStorage.deployed();
    });

    it("should deploy with the correct owner", async function () {
        expect(postStorage.address).to.exist;
    });

    it("should allow users to post and emit an event", async function () {
        const title = "First Post";
        const content = "This is my first blog post!";
        const category = 0; // Category.Blog

        await expect(postStorage.post(title, content, category, addr1.address))
            .to.emit(postStorage, "PostPublished")
            .withArgs(addr1.address, title, category);

        const posts = await postStorage.getUsersPost(addr1.address);
        expect(posts.length).to.equal(1);
        expect(posts[0].title).to.equal(title);
        expect(posts[0].content).to.equal(content);
        expect(posts[0].author).to.equal(addr1.address);
        expect(posts[0].category).to.equal(category);
    });

    it("should allow a different caller to post on behalf of another user", async function () {
        const title = "Second Post";
        const content = "This post is created by addr2 but called by addr1.";
        const category = 1; // Category.News

        await postStorage.connect(addr1).post(title, content, category, addr2.address);

        const posts = await postStorage.getUsersPost(addr2.address);
        expect(posts.length).to.equal(1);
        expect(posts[0].title).to.equal(title);
        expect(posts[0].content).to.equal(content);
        expect(posts[0].author).to.equal(addr2.address);
        expect(posts[0].category).to.equal(category);
    });

    it("should retrieve posts by user address", async function () {
        const title1 = "First Post";
        const content1 = "This is my first blog post!";
        const category1 = 0; // Category.Blog

        const title2 = "Second Post";
        const content2 = "This is my second blog post!";
        const category2 = 1; // Category.News

        await postStorage.post(title1, content1, category1, addr1.address);
        await postStorage.post(title2, content2, category2, addr1.address);

        const posts = await postStorage.getUsersPost(addr1.address);
        expect(posts.length).to.equal(2);
        expect(posts[0].title).to.equal(title1);
        expect(posts[1].title).to.equal(title2);
    });

    it("should retrieve post by index", async function () {
        const title = "Indexed Post";
        const content = "This post can be retrieved by index!";
        const category = 0; // Category.Blog

        await postStorage.post(title, content, category, addr1.address);
        
        const post = await postStorage.getPostByIndex(addr1.address, 0);
        expect(post.title).to.equal(title);
        expect(post.content).to.equal(content);
        expect(post.author).to.equal(addr1.address);
    });

    it("should revert when accessing a non-existing post by index", async function () {
        await expect(postStorage.getPostByIndex(addr1.address, 0))
            .to.be.revertedWith("Post does not exist");
    });

    it("should retrieve all posts", async function () {
        const title1 = "First Post";
        const content1 = "This is my first blog post!";
        const category1 = 0; // Category.Blog

        const title2 = "Second Post";
        const content2 = "This is my second blog post!";
        const category2 = 1; // Category.News

        await postStorage.post(title1, content1, category1, addr1.address);
        await postStorage.post(title2, content2, category2, addr1.address);
        await postStorage.post("Another Post", "Content", 2, addr2.address);

        const allPosts = await postStorage.getAllPosts();
        expect(allPosts.length).to.equal(3);
        expect(allPosts[0].title).to.equal(title1);
        expect(allPosts[1].title).to.equal(title2);
        expect(allPosts[2].author).to.equal(addr2.address);
    });

    it("should retrieve all users", async function () {
        await postStorage.post("Post 1", "Content 1", 0, addr1.address);
        await postStorage.post("Post 2", "Content 2", 1, addr2.address);

        const users = await postStorage.getAllUsers();
        expect(users.length).to.equal(2);
        expect(users).to.include(addr1.address);
        expect(users).to.include(addr2.address);
    });
});
