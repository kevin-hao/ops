import { logout } from "@/api/user";
import { login, getUserInfo, getMoveRouter } from "@/api/index";
import { getToken, setToken, removeToken } from "@/utils/auth";
import { resetRouter } from "@/router";

const getDefaultState = () => {
  return {
    token: getToken(),
    name: "",
    avatar: "",
    menus: "",  //存放路由表的容器
  };
};

const state = getDefaultState();

const mutations = {
  RESET_STATE: state => {
    Object.assign(state, getDefaultState());
  },
  SET_TOKEN: (state, token) => {
    state.token = token;
  },
  SET_NAME: (state, name) => {
    state.name = name;
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar;
  },
  SET_MENU: (state, menus) => {
    state.menus = menus;
  }
};

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { username, password } = userInfo;
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password }).then(response => {
          //const { data } = response;
          commit("SET_TOKEN", response.token);
          setToken(response.token);
          resolve();
        })
        .catch(error => {
          reject(error);
        });
    });
  },

  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getUserInfo().then(response => {
          const { data } = response;
          if (!data) {
            return reject("Verification failed, please Login again.");
          }
          const { name, avatar } = data;

          commit("SET_NAME", name);
          commit("SET_AVATAR", avatar);
          resolve(data);
        })
        .catch(error => {
          reject(error);
        });
    });
  },

  // get router
  getRouter({ commit, state }) {
    return new Promise((resolve, reject) => {
      getMoveRouter().then(response => {
          const menus = response.data;
          //如果需要404页面，在这里添加
          menus.push(
          {
            path: "/404",
            component: "404",
            hidden: true
          },
          {
            path: "*",
            component: "404",
            hidden: true
          }
          )
          commit("SET_MENU", menus);
          resolve();
        })
        .catch(error => {
          reject(error);
        });
    });
  },

  // user logout
  //   logout({ commit, state }) {
  //     return new Promise((resolve, reject) => {
  //       logout(state.token)
  //         .then(() => {
  //           removeToken(); // must remove  token  first
  //           resetRouter();
  //           commit("RESET_STATE");
  //           resolve();
  //         })
  //         .catch(error => {
  //           reject(error);
  //         });
  //     });
  //   },
  logout({ commit, state }) {
    return new Promise((resolve, reject) => {
      logout().then(() => {
          removeToken(); // must remove  token  first
          resetRouter();
          commit("RESET_STATE");
          resolve();
        })
        .catch(error => {
          reject(error);
        });
    });
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      removeToken(); // must remove  token  first
      commit("RESET_STATE");
      resolve();
    });
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
