import { useRecoilValue, useSetRecoilState } from "recoil";
import {userState} from "../store/atoms"
import { useNavigate } from "react-router-dom";
import { authService } from "../services/authServices";

export default function DashboardPage() {
    const user = useRecoilValue(userState)
    const setUser = useSetRecoilState(userState)
    const navigate = useNavigate();

    const handleLogout = async () => {
        await authService.logout();
        setUser(null)
        navigate('/login')
    };

  return (
    <>
      <div className="min-h-screen bg-gray-100">
        <header className="bg-white shadow">
          <div className="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8 flex justify-between items-center">
            <h1 className="text-3xl font-bold text-gray-900">Dashboard</h1>
            <button
              onClick={handleLogout}
              className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 cursor-pointer"
            >
              Sign Out
            </button>
          </div>
        </header>

        <main>
            <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
                <div className="px-4 py-6 sm:px-0">
                    <div className="border-4 border-dashed border-gray-200 rounded-lg p-8 text-center">
                        {user ? (
                            <div>
                                <h2 className="text-2xl font-bold mb-4">Welcome, {user.name || 'User'}</h2>
                                <p className="text-gray-600">You are now logged in.</p>
                            </div>
                        ):(
                            <p className="text-gray-600">Loading user data...</p>
                        )}
                    </div>
                </div>
            </div>
        </main>
      </div>
    </>
  );
}
