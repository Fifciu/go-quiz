import { Cookies } from "quasar";
import { RouterMiddlewareParam } from "../types";

export default ({ next, router }: RouterMiddlewareParam) => {
  if (Cookies.get(import.meta.env.cookie_token_key)) {
    return router.push({ 
      name: 'user-dashboard'
    })
  }
  next();
};
