<template>
  <el-table :data="users">
    <el-table-column prop="name" label="用户名" width="">
    </el-table-column>

    <el-table-column prop="gender" :formatter="getGender" label="性别">
    </el-table-column>
    <el-table-column prop="email" label="邮箱" width="">
    </el-table-column>
    <el-table-column prop="created_at" label="加入时间" :formatter="dateFormat" width="">

    </el-table-column>
  </el-table>
</template>

<script>
  import moment from 'moment'
  export default {
      name: "Home",
      data(){
        return {
          users: []
        }
      },
      created() {
        var that = this
        this.axios({
          method:"get",
          url: "http://localhost:8080/user/query"
        }).then(function(res){
          that.users = res.data.msg
          console.log(that.users)
        })
      },
      methods:{
        dateFormat:function(row,column){
          var date = row[column.property];
          if(date == undefined){return ''};
          return moment(date).format("YYYY-MM-DD HH:mm:ss")
        },
        getGender:function(row,column){
         var gender = row[column.property]
         if(gender == undefined){
           return ''
         } else if (gender == 1) {
           return ' 男'
         } else if (gender == 0){
           return '女'
         } else {
           return '未知'
         }
        }
      }

  }
</script>

<style>
</style>
