import {apiClient} from "./client.js";
import {jwtDecode} from "jwt-decode";

export const ACCESS_TOKEN_KEY = "accessToken";

export const setAccessToken = token => {
    localStorage.setItem(ACCESS_TOKEN_KEY, token);
};

export const getAccessToken = () => {
    return localStorage.getItem(ACCESS_TOKEN_KEY);
};

export const clearAccessToken = () => {
    localStorage.removeItem(ACCESS_TOKEN_KEY);
};

apiClient.interceptors.request.use((config) => {
        const token = getAccessToken();
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => Promise.reject(error));

apiClient.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalRequest = error.config;

        if (error.response?.status === 401 && !originalRequest._retry) {
            originalRequest._retry = true;

            if (originalRequest.url.includes("/auth/login") || originalRequest.url.includes("/auth/register") || originalRequest.url.includes("/auth/refresh")) {
                return Promise.reject(error);
            }

            try {
                const res = await apiClient.post("/auth/refresh");
                const newToken = res.data.access_token;

                setAccessToken(newToken);
                originalRequest.headers.Authorization = `Bearer ${newToken}`;

                return apiClient(originalRequest);
            } catch {
                await apiClient.post("/auth/logout");
                clearAccessToken();
                window.location.href = '/login';
                return Promise.reject(error);
            }
        }

        return Promise.reject(error);
    }
);

export const decodeUser = token => {
    const decoded = jwtDecode(token);
    return {
        role: decoded.role
    };
};

export const login = async (data) => {
    const res = await apiClient.post("/auth/login", data);
    setAccessToken(res.data.access_token);
    return decodeUser(res.data.access_token);
};

export const logout = async () => {
    await apiClient.post("/auth/logout");
    clearAccessToken();
};
