import { useRecoilValue } from 'recoil';
import { Navigate, Outlet } from 'react-router-dom';
import { userState } from '../store/atoms';

export default function ProtectedRoute() {
    const user = useRecoilValue(userState);

    // If user is not logged in, redirect to login page
    if (!user) {
        return <Navigate to="/login" replace />;
    }

    // Otherwise, render the child routes
    return <Outlet />;
} 