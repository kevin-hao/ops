<template>
    <div>
      <el-container style="height: 500px; border: 1px solid #eee">
        <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
          <el-menu :default-openeds="['1']">
              <el-submenu index="1">
                  <template slot="title">用户管理</template>
                  <el-menu-item-group>
                    <el-menu-item index="1-1"><router-link to="/userlist" >用户列表</router-link></el-menu-item>
                    <el-menu-item index="1-2"><router-link :to="{name:'Userinfo', params:{id: 3}}">用户信息</router-link></el-menu-item>
                  </el-menu-item-group>
                </el-submenu>
          </el-menu>
        </el-aside>
        <el-container>
            <el-header style="text-align: right; font-size: 12px">
              <el-dropdown>
                <i class="el-icon-setting" style="margin-right: 15px"></i>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item>查看</el-dropdown-item>
                  <el-dropdown-item><router-link to="/login">退出</router-link></el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
              <span>管理员</span>
            </el-header>

            <el-main>
              <router-view></router-view>
            </el-main>
          </el-container>
      </el-container>
    </div>
</template>
<style>
  .el-header {
    background-color: #B3C0D1;
    color: #333;
    line-height: 60px;
  }

  .el-aside {
    color: #333;
  }
</style>
<script>
export default {
    methods:{
      logout:function(){
        this.$route.push("/login")
      }
    }
}
</script>
