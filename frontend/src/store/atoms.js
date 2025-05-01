import { atom } from "recoil";
import { authService } from "../services/authServices";


export const userState = atom({
    key: 'userState',
    default: null,
});

export const authLoadingState = atom({
    key:'authLoadingState',
    default: false,
});

export const authErrorState = atom({
    key:'authErrorState',
    default: null,
});

export const authState = atom({
    key: 'authState',
    default :{
        authAccessToken : localStorage.getItem('access_token') || null,
        authRefreshToken : localStorage.getItem('refresh_token') || null,
        isAuthenticated : authService.isAuthenticated()
    }
});