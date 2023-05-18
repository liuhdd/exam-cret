/**
 * 登录请求参数
 */
export interface LoginData {
    /**
     * 用户名
     */
    username: string;
    /**
     * 密码
     */
    password: string;
  }
  
  /**
   * 登录响应
   */
  export interface LoginResult {
    /**
     * 访问token
     */
    token?: string;
    /**
     * 过期时间(单位：毫秒)
     */
    expires?: number;
    /**
     * 刷新token
     */
    refreshToken?: string;
    /**
     * token 类型
     */
    tokenType?: string;

    role:string;
  }
  
  export interface User {
    user_id: string;
    username: string;
    password: string;
    role: string;
  }