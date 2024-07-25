import axios from '@/axios';
import { getToken } from '@/utils/auth';
import { ALL_POSTS_URL, GET_POST_BY_ID_UEL as GET_POST_BY_ID_URL, USER_POSTS_URL } from '@/utils/request_paths';
import { mapPost, mapPosts } from '@/scripts/post_model'


export const getUserPosts = async (pageNumber, pageSize) => {
    const token = getToken();
    try {
        const response = await axios.get(USER_POSTS_URL, {
            headers: {
                Authorization: `Bearer ${token}`
            },
            params: {
                page: pageNumber,
                pageSize: pageSize
            }
        });

        const posts = mapPosts(response.data)

        return posts;

    } catch (error) {
        console.error('Error fetching user posts:', error);
        return [];
    }
};

export const getAllPosts = async (pageNumber, pageSize) => {
    try {
        const response = await axios.get(ALL_POSTS_URL, {
            params: {
                page: pageNumber,
                pageSize: pageSize
            }
        });

        const posts = mapPosts(response.data)

        return posts;

    } catch (error) {
        console.error('Error fetching posts:', error);
        return [];
    }

};

export const getPostById = async (postId) => {
    try {
        const response = await axios.get(GET_POST_BY_ID_URL + postId, {
        });

        const posts = mapPost(response.data)

        return posts;

    } catch (error) {
        console.error('Error fetching posts:', error);
        return [];
    }

};