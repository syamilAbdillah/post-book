import { UserService, PostService } from './schema.gen'

export const userService = new UserService('', window.fetch)
export const postService = new PostService('', window.fetch)