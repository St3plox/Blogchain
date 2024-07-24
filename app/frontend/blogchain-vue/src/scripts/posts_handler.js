import axios from '@/axios';
import { getToken } from '@/utils/auth';
import { USER_POSTS_URL } from '@/utils/request_paths';
import { mapPosts } from '@/scripts/post_model'


const getUserPosts = async (pageNumber, pageSize) => {
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

        console.log(response.data)
        const posts = mapPosts(response.data)
        console.log(posts)

        return posts;

    } catch (error) {
        console.error('Error fetching user posts:', error);
        return [];
    }
};

export default getUserPosts;
