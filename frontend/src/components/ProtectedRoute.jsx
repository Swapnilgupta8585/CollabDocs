import { useRecoilValue, useSetRecoilState } from 'recoil';
import { Navigate, Outlet } from 'react-router-dom';
import { userState, authState } from '../store/atoms';
import { useEffect } from 'react';
import { authService } from '../services/authServices';


export default function ProtectedRoute() {
    const user = useRecoilValue(userState);
    const auth = useRecoilValue(authState);
    const setAuthState = useSetRecoilState(authState)


    useEffect(() => {
        //update auth state on mount to ensure latest token store
        setAuthState({
            authAccessToken: authService.getAccessToken(),
            authRefreshToken: authService.getRefreshToken(),
            isAuthenticated: authService.isAuthenticated()
        });
    }, [setAuthState]);


    // If user is not logged in or there's no valid token, redirect to login page
    if (!user || !auth.isAuthenticated) {
        return <Navigate to="/login" replace />;
    }


    // Otherwise, render the child routes
    return <Outlet />;
} 