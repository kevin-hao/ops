import request from "@/utils/request";

//登录接口
export function login(data) {
  return request({
    url: "/auth/login/",
    method: "post",
    data
  });
}

// 获取用户详情
export function getUserInfo(data) {
  return request({
    url: "/user/getuserinfo",
    method: "post",
    data
  });
}
//退出
export function logout() {
  return request({
    url: "/auth/logout",
    method: "post"
  });
}
