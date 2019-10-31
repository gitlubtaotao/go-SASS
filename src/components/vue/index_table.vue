<template>
    <table :id="options.id" class="table table-bordered table-hover">
        <thead>
        <tr>
            <th v-if="showActions()">操作</th>
            <th v-for="(item,index) in colNames" :class="item.class" v-if="item.key !=='Id'">{{item.value}}</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(record,index) in objects">
            <td v-if="showActions()">
                <button v-if="isShowButton('show')" class="btn btn-xs btn-success"
                        @click="showMethod(record['Id'],index)">
                    <i class="ace-icon fa fa-check bigger-120"></i>
                    详情
                </button>
                <button v-if="isShowButton('edit')" class="btn btn-xs btn-info" @click="editMethod(record['Id'],index)">
                    <i class="ace-icon fa fa-pencil bigger-120"></i>
                    修改
                </button>
                <button v-if="isShowButton('destroy')" class="btn btn-xs btn-danger"
                        @click="destroyMethod(record['Id'],index)">
                    <i class="ace-icon fa fa-trash-o bigger-120"></i>
                    删除
                </button>
            </td>
            <td v-for="item in colNames" v-if="item.key !== 'Id'">{{showItem(record,item.key)}}</td>
        </tr>
        </tbody>
    </table>
</template>
<script>
    export default {
        props: {
            colNames: {
                type: Array,
                required: true
            },
            objects: {
                type: Array,
                required: true,
            },
            options: {
                type: Object,
                default: function () {
                    return {id: 'index-table', class: ''}
                }
            },
            actions: {}
        },
        computed: {},
        methods: {
            isShowButton: function (action) {
                let value = false;
                if (this.options[action] !== "" && typeof (this.actions[action]) !== "undefined") {
                    value = true;
                }
                return value;
            },
            //是否显示操作
            showActions: function () {
                let value = false;
                if (typeof (this.actions) !== 'undefined') {
                    value = true;
                }
                return value;
            },

            //取数据的表关联的其他字段的值
            showItem: function (record, item) {
                let arrayItem = item.split('.');
                if (record[item] !== "" && typeof (record[item]) !== 'undefined') {
                    return record[item]
                }
                if (arrayItem.length === 2) {
                    let value = record[arrayItem[0]];
                    if (typeof (value) === 'undefined' || value === "") {
                        return ""
                    } else {
                        return value[arrayItem[1]];
                    }
                }
            },
            editMethod: function (id, index) {
                let url = this.actions.edit;
                url = url.replace(":id",id);
                window.location.href = url;
            },
            showMethod: function (id, index) {

            },
            //删除对应的记录
            destroyMethod: function (id, index) {
                let _this = this;
                let url = this.actions.destroy.replace(":id", id);
                if (confirm("确定删除该记录？")) {
                    axios.delete(url, {
                        headers: {'X-Requested-With': 'XMLHttpRequest'},
                        dataType: 'json',
                    }).then(function (response) {
                        console.log(response);
                        if (response.data === "OK") {
                            toastr.success("删除成功");
                            _this.objects.splice(index, 1)
                        } else {
                            toastr.error(response.data);
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                }
            }
        },
    };
</script>
<style lang="css" scoped>
    a {
        cursor: pointer;
    }
</style>