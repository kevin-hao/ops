<template>
    <el-container>
        <el-header>
            <el-menu :default-active="activeIndex2" class="el-menu-demo" mode="horizontal" @select="handleSelect" background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
                <el-menu-item index="1">首页</el-menu-item>
                <el-submenu index="2">
                    <template slot="title">用户管理</template>
                    <el-menu-item index="2-1">用户列表</el-menu-item>
                </el-submenu>
            </el-menu>
        </el-header>
        <el-container>
            <el-main></el-main>
        </el-container>
    </el-container>
</template>


<script>
  export default {
    data() {
      return {
        activeIndex: '1',
        activeIndex2: '1'
      };
    },
    methods: {
      handleSelect(key, keyPath) {
        console.log(key, keyPath);
      }
    }
  }
</script>