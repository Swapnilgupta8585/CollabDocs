import apiClient from "./apiClient";

const API_URL = "http://localhost:8080/api";

export const authService = {
  async register(name, email, password) {
    try {
      const response = await apiClient.post(`${API_URL}/users`, {
        name,
        email,
        password,
      });
      return response.data.user;

    } catch (error) {
      throw new Error(
        error.response?.data?.message || error.message || "Something went wrong"
      );
    }
  },

  async login(email, password) {
    try {
      const response = await apiClient.post(`${API_URL}/login`, {
        email,
        password,
      });

      if (response.data.token && response.data.refresh_token) {
        this.setTokens(response.data.token, response.data.refresh_token)
      }

      return response.data.user;
    } catch (error) {
      throw new Error(
        error.response?.data?.message || error.message || "Something went wrong"
      );
    }
  },

  async logout() {
    try {
      await apiClient.post(`${API_URL}/logout`);
      this.clearTokens();
      return true;
    } catch (error) {
      console.error("Logout error:", error);
      return false;
    }
  },

  // token management helpers
  setTokens(accessToken, refreshToken){
    localStorage.setItem('access_token', accessToken);
    localStorage.setItem('refresh_token', refreshToken);
  },

  clearTokens(){
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');
  },

  getAccessToken(){
    return localStorage.getItem('access_token');
  },

  getRefreshToken(){
    return localStorage.getItem('refresh_token');
  },

  isAuthenticated() {
    return !!localStorage.getItem('access_token')
  }

};

