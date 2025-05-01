import axios from "axios";


const API_URL = "http://localhost:8080/api";

// create axios instance with some configuration
const apiClient = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// request interceptors for adding token to headers
apiClient.interceptors.request.use(
  (config) => {
    const accessToken = localStorage.getItem("access_token");
    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// response interceptor for handling token refresh
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    // if error is 401 Unauthorized and we haven't already tried to refresh
    if (
      error.response &&
      error.response.status === 401 &&
      !originalRequest._retry
    ) {
      originalRequest._retry = true;

      try {
        //attempt to refresh the token
        const refreshToken = localStorage.getItem("refresh_token");
        if (!refreshToken) {
          // no refresh token, redirect to login
          return Promise.reject(error);
        }

        // call refresh endpoint
        const response = await axios.post(
          `${API_URL}/refresh`,
          {},
          {
            headers: {
              Authorization: `Bearer ${refreshToken}`,
            },
          }
        );

        // update the token in the localstorage
        const { token } = response.data;
        localStorage.setItem("access_token", token);

        //update the auth header for the original request
        originalRequest.headers.Authorization = `Bearer ${token}`;

        // retry the original request
        return apiClient(originalRequest);
      } catch (refreshError) {
        // refresh failed - clear tokens and redirect to Login
        localStorage.removeItem("access_token");
        localStorage.removeItem("refresh_token");

        return Promise.reject(refreshError);
      }
    }
    return Promise.reject(error);
  }
);

export default apiClient