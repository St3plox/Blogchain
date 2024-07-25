export const Category = {
  BLOG: 0,
  NEWS: 1,
  ARTICLE: 2,
};

//post_model.js
export const categoryNames = {
  [Category.BLOG]: 'Blog',
  [Category.NEWS]: 'News',
  [Category.ARTICLE]: 'Article'
};

export class Post {
  /**
   * Create a Post.
   * @param {BigInt} id - The unique identifier for the post.
   * @param {string} author - The address of the author.
   * @param {string} title - The title of the post.
   * @param {string} content - The content of the post.
   * @param {BigInt} timestamp - The timestamp of when the post was created.
   * @param {number} category - The category of the post.
   * @param {string[]} tags - The tags associated with the post.
  */
  constructor(id, author, title, content, timestamp, category, tags) {
    this.id = id;
    this.author = author;
    this.title = title;
    this.content = content;
    this.timestamp = timestamp;
    this.category = category;
    this.tags = tags;
  }
}

export class NewPost {

  /**
   * Create a NewPost.
   * @param {string} title - The title of the new post.
   * @param {string} content - The content of the new post.
   * @param {number} category - The category of the new post.
   * @param {string[]} tags - The tags associated with the new post.
  */

  constructor(title, content, category, tags) {
    this.title = title;
    this.content = content;
    this.category = category;
    this.tags = tags;
  }
}

export const mapPosts = (postsData) => {
  return postsData.map(postData => new Post(
    postData.id,
    postData.author,
    postData.title,
    postData.content,
    postData.timestamp,
    postData.category,
    postData.tags
  ));
};

export const mapPost = (postData) => {
  return new Post(
    postData.id,
    postData.author,
    postData.title,
    postData.content,
    postData.timestamp,
    postData.category,
    postData.tags
  )
}

