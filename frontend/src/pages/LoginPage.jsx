import { useRecoilValue } from 'recoil';
import { Navigate } from 'react-router-dom';
import { userState, authLoadingState, authErrorState } from '../store/atoms';
import LoginForm from '../components/LoginForm';

export default function LoginPage() {
    const user = useRecoilValue(userState);
    const loading = useRecoilValue(authLoadingState);
    const error = useRecoilValue(authErrorState);

    // Redirect if user is already logged in
    if (user) {
        return <Navigate to="/dashboard" replace />;
    }

    return (
        <div className="min-h-screen bg-gray-100 flex flex-col justify-center items-center py-12 sm:px-6 lg:px-8">
            <div className="sm:mx-auto sm:w-full sm:max-w-md">
                <h1 className="text-3xl font-extrabold text-center text-gray-900">
                    Welcome Back
                </h1>
                <p className="mt-2 text-sm text-center text-gray-600">
                    Sign in to access your account
                </p>
            </div>

            <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
                {error && (
                    <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-md mb-4" role="alert">
                        <p className="text-sm">{error}</p>
                    </div>
                )}

                {loading ? (
                    <div className="flex justify-center">
                        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
                    </div>
                ) : (
                    <LoginForm />
                )}
            </div>
        </div>
    );
} 