<template>
    <table :id="options.id" class="table table-bordered table-hover table-responsive">
        <thead>
        <tr>
            <th v-if="showActions()">操作</th>
            <th v-for="(item,index) in colNames" :class="item.class" v-if="item.key !=='Id'">{{item.value}}</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(record,index) in objects">
            <td v-if="showActions()" class="col-sm-2">
                <a href="javascript:void(0);" class="btn-xs btn-success btn btn-white"
                   @click="clickMethod(actions[0],record.Id,index)">{{actions[0]['name']}}</a>
                <div v-if="otherActions.length > 0" class="btn-group">
                    <button data-toggle="dropdown" class="btn btn-sm btn-info dropdown-toggle btn-white">
                        更多
                        <i class="ace-icon fa fa-angle-down icon-on-right"></i>
                    </button>
                    <ul class="dropdown-menu">
                        <li v-for="action in otherActions">
                            <a href="javascript:void(0);" @click="clickMethod(action,record.Id,index)">{{action['name']}}</a>
                        </li>
                    </ul>
                </div>
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
            actions: {
                type: Array,
                required: true
            }
        },
        computed: {
            //其他actions
            otherActions: function () {
                let other_actions = this.actions;
                return other_actions.slice(1, this.actions.length);
            },
        },
        methods: {
            //是否显示操作
            showActions: function () {
                let value = true;
                if (this.actions.length === 0) {
                    value = false;
                }
                return value;
            },
            //取数据的表关联的其他字段的值
            showItem: function (record, item) {
                if (toString.call(record[item]) === '[object Object]') {
                    return record[item]['Name']
                }
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
            clickMethod: function (action, id, index) {
                let url = action['url'];
                url = url.replace(":id", id);
                if (action.method === 'delete') {
                    this.destroyMethod(url, index)
                } else if (action['remote'] === false) {
                    window.location.href = url;
                } else {

                }
            },
            //删除对应的记录
            destroyMethod: function (url, index) {
                let _this = this;
                if (confirm("确定删除该记录？")) {
                    axios.delete(url, {
                        headers: {'X-Requested-With': 'XMLHttpRequest'},
                        dataType: 'json',
                    }).then(function (response) {
                        console.log(response);
                        if (response.data.code === 200) {
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
    td,th{
        font-size: 12px;
    }

    .dropdown-menu {
        min-width: 65px;
        padding: 2px;
    }
</style>