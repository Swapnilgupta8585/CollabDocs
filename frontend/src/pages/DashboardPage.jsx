import { useRecoilValue, useSetRecoilState } from "recoil";
import { userState, authState } from "../store/atoms";
import { useNavigate } from "react-router-dom";
import { authService } from "../services/authServices";

export default function DashboardPage() {
  const user = useRecoilValue(userState);
  const setUser = useSetRecoilState(userState);
  const navigate = useNavigate();
  const setAuth = useSetRecoilState(authState)

  const handleLogout = async () => {
    try {
      await authService.logout();
      // Clear all state in a single update
      setAuth({
        authAccessToken: null,
        authRefreshToken: null,
        isAuthenticated: false,
      });
      setUser(null);
      // Navigate after state is cleared
      navigate("/login", { replace: true });
    } catch (error) {
      console.error("Logout error:", error);
    }
  };

  return (
    <>
      <>
        <div className="min-h-screen bg-gray-50">
          <header className=" shadow-sm bg-amber-50">
            <div className="max-w-7xl mx-auto py-4 px-6 flex justify-between items-center">
              <h1 className="text-2xl font-semibold text-gray-800">
                Dashboard
              </h1>
              <button
                onClick={handleLogout}
                className="px-4 py-2 text-sm font-medium rounded-lg text-white bg-red-600 transition focus:outline-none focus:ring-2 focus:ring-red-400"
              >
                Sign Out
              </button>
            </div>
          </header>

          <main>
            <div className="max-w-7xl mx-auto py-8 px-6 ">
              <div className="rounded-lg border border-gray-200  bg-amber-50 p-8 text-center shadow-sm">
                {user ? (
                  <div>
                    <h2 className="text-xl font-semibold text-gray-800 mb-2">
                      Welcome, {user.name || "User"}
                    </h2>
                    <p className="text-gray-500">You're now logged in.</p>
                  </div>
                ) : (
                  <p className="text-gray-500">Loading user data...</p>
                )}

                <div className="mt-6 mb-8 max-w-sm mx-auto">
                  <button className="w-full px-4 py-3 rounded-lg bg-green-600 hover:bg-green-700 text-white text-sm font-medium transition focus:outline-none focus:ring-2 focus:ring-green-400 cursor-pointer">
                    + Create Doc
                  </button>
                </div>

                <div className="flex flex-wrap justify-center gap-6">
                  {Array.from({ length: 16 }).map((_, i) => (
                    <div
                      key={i}
                      className="aspect-[6/5] w-full max-w-[15rem] bg-white text-left border border-gray-200 rounded-xl shadow-sm px-4 py-5 flex flex-col justify-between hover:shadow-lg transition cursor-pointer"
                    >
                      <div className="text-sm font-medium text-gray-800 mb-1">
                        Document #{i + 1}
                      </div>
                      <div className="text-xs text-gray-500">
                        Last updated: 5 min ago
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          </main>
        </div>
      </>
    </>
  );
}
