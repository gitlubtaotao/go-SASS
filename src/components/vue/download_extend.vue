<template>
    <div>
        <p>{{$t('download')}}</p>
        <ul>
            <li>
                <a href="javascript:void(0);"
                   class="btn btn-white btn-warning btn-bold btn-xs"
                   @click="downloadXls()"
                >XLS</a>
            </li>
            <li>
                <a href="javascript:void(0);"
                   class="btn btn-white btn-default btn-xs"
                   @click="downloadPdf()"
                >PDF</a>
            </li>
        </ul>
    </div>
</template>
<script>
    function getCookie(cname) {
        let name = cname + "=";
        let ca = document.cookie.split(';');
        for (let i = 0; i < ca.length; i++) {
            let c = ca[i];
            while (c.charAt(0) === ' ') {
                c = c.substring(1);
            }
            if (c.indexOf(name) === 0) {
                return c.substring(name.length, c.length);
            }
        }
        return "";
    }

    export default {
        props: {
            page: {
                type: Number,
                default: 1
            },
            clickHandler: {
                type: String,
                default: ''
            }
        },
        i18n: {
            locale: getCookie('lang'),
            messages: {
                'zh-CN': {download: '更多'},
                'en-US': {download: 'Download'},
            }
        },
        data() {
            return {
                checked: false,
                title: 'Check me'
            }
        },
        methods: {
            downloadXls: function () {
                let url = location.pathname.replace(/\/index/, "");
                url += ("?query=" + this.clickHandler + "&page=" + this.page + "&format=xls");
                window.location = url;
            },
            downloadPdf: function () {
            }
        }
    }
</script>
<style lang="css" scoped>
    p {
        display: inline-block;
        color: #70a0c1 !important;
    }

    ul {
        display: inline-block;
        margin: 0 0 0 5px;
    }

    ul > li {
        display: inline;
        margin-left: 5px;
    }
</style>