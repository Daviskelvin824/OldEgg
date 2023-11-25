import authenticate from "../Pages/api-calls/auth/authenticate";
import getCookie from "../utility/getCookie";
import { useEffect, useState } from "react";

const useAuth = () => {

  const [user, setUser] = useState({});

  useEffect(() => {

    const getUser = async () => {

      const cookie = getCookie("Auth")
      const result = await authenticate({
        "token_string": cookie
      });
      if (result.role_id) setUser(result);
      else if (result.shop_name)
        setUser({
          id: result.ID,
          first_name: result.shop_name,
          email: result.shop_email,
          password: result.shop_password,
          role_id: 3,
          status: result.status
        })
      else{
        setUser({})
      }
  
    }
  
    getUser();

  }, [])

  return user

}

export default useAuth;
