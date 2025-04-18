import axios from "axios";

const API_URL = "http://localhost:8080/api";

export const authService = {
  async register(email, password) {
    try {
      const response = await axios.post(`${API_URL}/users`, {
        name,
        email,
        password,
      });
      return response.data;
    } catch (error) {
      throw new Error(
        error.response?.data?.message || error.message || "Something went wrong"
      );
    }
  },

  async login(email, password) {
    try {
      const response = await axios.post(`${API_URL}/login`, {
        email,
        password,
      });
      return response.data;
    } catch (error) {
      throw new Error(
        error.response?.data?.message || error.message || "Something went wrong"
      );
    }
  },

  async logout() {
    try {
      await axios.post(`${API_URL}/logout`);
      return true;
    } catch (error) {
      console.error("Logout error:", error);
      return false;
    }
  },
};
