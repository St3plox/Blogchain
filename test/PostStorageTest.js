const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("PostStorage", function () {
    let PostStorage, postStorage, owner, addr1;
    
    beforeEach(async function () {
        PostStorage = await ethers.getContractFactory("PostStorage");
        [owner, addr1, ...addrs] = await ethers.getSigners();
        postStorage = await PostStorage.deploy();
        await postStorage.deployed();
    });

    it("Should allow a user to create a post", async function () {
        const title = "My first post";
        const content = "Hello, this is my first post!";
        const category = 0; // Category.Blog

        await expect(postStorage.post(title, content, category))
            .to.emit(postStorage, 'PostPublished')
            .withArgs(owner.address, title, category);

        const myPosts = await postStorage.getMyPosts();
        expect(myPosts.length).to.equal(1);
        expect(myPosts[0].title).to.equal(title);
        expect(myPosts[0].content).to.equal(content);
        expect(myPosts[0].category).to.equal(category);
    });

    it("Should allow a user to retrieve their posts", async function () {
        const title1 = "Post 1";
        const content1 = "Content for post 1";
        const category1 = 1; // Category.News

        const title2 = "Post 2";
        const content2 = "Content for post 2";
        const category2 = 2; // Category.Article

        await postStorage.post(title1, content1, category1);
        await postStorage.post(title2, content2, category2);

        const myPosts = await postStorage.getMyPosts();
        expect(myPosts.length).to.equal(2);
        expect(myPosts[0].title).to.equal(title1);
        expect(myPosts[0].content).to.equal(content1);
        expect(myPosts[0].category).to.equal(category1);
        expect(myPosts[1].title).to.equal(title2);
        expect(myPosts[1].content).to.equal(content2);
        expect(myPosts[1].category).to.equal(category2);
    });

    it("Should allow retrieving another user's posts", async function () {
        const title = "Addr1's post";
        const content = "Content for addr1's post";
        const category = 1; // Category.News

        await postStorage.connect(addr1).post(title, content, category);

        const userPosts = await postStorage.getUsersPost(addr1.address);
        expect(userPosts.length).to.equal(1);
        expect(userPosts[0].title).to.equal(title);
        expect(userPosts[0].content).to.equal(content);
        expect(userPosts[0].category).to.equal(category);
    });
});
