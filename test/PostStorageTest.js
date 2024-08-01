const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("PostStorage", function () {
    let PostStorage;
    let postStorage;
    let owner;
    let addr1;
    let addr2;

    const title1 = "First Post";
    const content1 = "This is my first blog post!";
    const category1 = 0; // Category.Blog

    const title2 = "Second Post";
    const content2 = "This is my second blog post!";
    const category2 = 1; // Category.News

    const title3 = "Another Post";
    const content3 = "This post can be retrieved by index!";
    const category3 = 2; // Category.Article

    const mediaNames1 = ["image1", "video1"];
    const mediaUrls1 = ["https://example.com/image1.jpg", "https://example.com/video1.mp4"];

    const mediaNames2 = ["image2"];
    const mediaUrls2 = ["https://example.com/image2.jpg"];

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
        const tags = ["tag1", "tag2"];

        await expect(postStorage.post(title1, content1, category1, tags, addr1.address, mediaNames1, mediaUrls1))
            .to.emit(postStorage, "PostPublished")
            .withArgs(1, addr1.address, title1, tags, category1);

        const posts = await postStorage.getUsersPost(addr1.address);
        expect(posts.length).to.equal(1);
        expect(posts[0].title).to.equal(title1);
        expect(posts[0].content).to.equal(content1);
        expect(posts[0].author).to.equal(addr1.address);
        expect(posts[0].category).to.equal(category1);
        expect(posts[0].media_names).to.deep.equal(mediaNames1);
        expect(posts[0].media_urls).to.deep.equal(mediaUrls1);
    });

    it("should allow a different caller to post on behalf of another user", async function () {
        const tags = ["tag3"];

        await postStorage.connect(addr1).post(title2, content2, category2, tags, addr2.address, mediaNames2, mediaUrls2);

        const posts = await postStorage.getUsersPost(addr2.address);
        expect(posts.length).to.equal(1);
        expect(posts[0].title).to.equal(title2);
        expect(posts[0].content).to.equal(content2);
        expect(posts[0].author).to.equal(addr2.address);
        expect(posts[0].category).to.equal(category2);
        expect(posts[0].media_names).to.deep.equal(mediaNames2);
        expect(posts[0].media_urls).to.deep.equal(mediaUrls2);
    });

    it("should retrieve posts by user address", async function () {
        const tags = ["tag1", "tag2"];

        await postStorage.post(title1, content1, category1, tags, addr1.address, mediaNames1, mediaUrls1);
        await postStorage.post(title2, content2, category2, tags, addr1.address, mediaNames2, mediaUrls2);

        const posts = await postStorage.getUsersPost(addr1.address);
        expect(posts.length).to.equal(2);
        expect(posts[0].title).to.equal(title1);
        expect(posts[1].title).to.equal(title2);
    });

    it("should retrieve post by index", async function () {
        const tags = ["tag1"];

        await postStorage.post(title3, content3, category3, tags, addr1.address, mediaNames1, mediaUrls1);
        
        const post = await postStorage.getPostByIndex(addr1.address, 0);
        expect(post.title).to.equal(title3);
        expect(post.content).to.equal(content3);
        expect(post.author).to.equal(addr1.address);
        expect(post.media_names).to.deep.equal(mediaNames1);
        expect(post.media_urls).to.deep.equal(mediaUrls1);
    });

    it("should revert when accessing a non-existing post by index", async function () {
        await expect(postStorage.getPostByIndex(addr1.address, 0))
            .to.be.revertedWith("Post does not exist");
    });

    it("should retrieve all posts", async function () {
        const tags = ["tag1", "tag2"];

        await postStorage.post(title1, content1, category1, tags, addr1.address, mediaNames1, mediaUrls1);
        await postStorage.post(title2, content2, category2, tags, addr1.address, mediaNames2, mediaUrls2);
        await postStorage.post("Another Post", "Content", 2, ["tag3"], addr2.address, mediaNames1, mediaUrls1);

        const allPosts = await postStorage.getAllPosts();
        expect(allPosts.length).to.equal(3);
        expect(allPosts[0].title).to.equal(title1);
        expect(allPosts[1].title).to.equal(title2);
        expect(allPosts[2].author).to.equal(addr2.address);
    });

    it("should retrieve all users", async function () {
        await postStorage.post("Post 1", "Content 1", 0, ["tag1"], addr1.address, mediaNames1, mediaUrls1);
        await postStorage.post("Post 2", "Content 2", 1, ["tag2"], addr2.address, mediaNames2, mediaUrls2);

        const users = await postStorage.getAllUsers();
        expect(users.length).to.equal(2);
        expect(users).to.include(addr1.address);
        expect(users).to.include(addr2.address);
    });

    it("should retrieve paginated posts by user", async function () {
        const tags = ["tag1"];
        
        await postStorage.post(title1, content1, category1, tags, addr1.address, mediaNames1, mediaUrls1);
        await postStorage.post(title2, content2, category2, tags, addr1.address, mediaNames2, mediaUrls2);
        await postStorage.post(title3, content3, category3, tags, addr1.address, mediaNames1, mediaUrls1);

        const pageSize = 2;
        const page1 = await postStorage.getUsersPostPaginated(addr1.address, 0, pageSize);
        expect(page1.length).to.equal(2);
        expect(page1[0].title).to.equal(title1);
        expect(page1[1].title).to.equal(title2);

        const page2 = await postStorage.getUsersPostPaginated(addr1.address, 1, pageSize);
        expect(page2.length).to.equal(1);
        expect(page2[0].title).to.equal(title3);
    });

    it("should revert when pagination is out of range", async function () {
        await postStorage.post(title1, content1, category1, ["tag1"], addr1.address, mediaNames1, mediaUrls1);

        await expect(postStorage.getUsersPostPaginated(addr1.address, 1, 1))
            .to.be.revertedWith("Page out of range");
    });

    it("should retrieve all posts paginated", async function () {
        const tags = ["tag1", "tag2"];

        await postStorage.post(title1, content1, category1, tags, addr1.address, mediaNames1, mediaUrls1);
        await postStorage.post(title2, content2, category2, tags, addr1.address, mediaNames2, mediaUrls2);
        await postStorage.post("Another Post", "Content", 2, ["tag3"], addr2.address, mediaNames1, mediaUrls1);

        const pageSize = 2;
        const page1 = await postStorage.getPostsPaginated(0, pageSize);
        expect(page1.length).to.equal(2);
        expect(page1[0].title).to.equal(title1);
        expect(page1[1].title).to.equal(title2);

        const page2 = await postStorage.getPostsPaginated(1, pageSize);
        expect(page2.length).to.equal(1);
        expect(page2[0].author).to.equal(addr2.address);
    });

    it("should revert when all posts pagination is out of range", async function () {
        const tags = ["tag1"];
        
        await postStorage.post(title1, content1, category1, tags, addr1.address, mediaNames1, mediaUrls1);

        await expect(postStorage.getPostsPaginated(1, 1))
            .to.be.revertedWith("Page out of range");
    });

    it("should retrieve post by ID", async function () {
        const tags = ["tag1"];
        
        await postStorage.post(title1, content1, category1, tags, addr1.address, mediaNames1, mediaUrls1);
        const post = await postStorage.getPostByID(1);
        expect(post.title).to.equal(title1);
        expect(post.content).to.equal(content1);
    });

    it("should revert when retrieving a non-existing post by ID", async function () {
        await expect(postStorage.getPostByID(1))
            .to.be.revertedWith("Post with this ID does not exist");
    });
});
