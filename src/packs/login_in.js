var app = new Vue({
    el: '#login_in',
    i18n: new VueI18n({
        locale: getCookie('lang'),
        messages: {
            "zh-CN": {blank: '%{name}不能为空', account: '账号', password: '密码'},
            "en-US": {blank: '%{name} Not Blank', account: 'Account', password: 'Password'}
        },
    }),
    data: {
        userName: '',
        Password: '',
        Remember: false,
        action: '/login'
    },
    methods: {
        submitForm: function () {
            if (!this.userName) {
                toastr.error(this.$i18n.t("blank",{name: this.$i18n.t('account')}));
                return false;
            }
            if (!this.Password) {
                toastr.error(this.$i18n.t("blank",{name: this.$i18n.t('password')}));
                return false;
            }
            axios({
                method: 'post',
                url: this.action,
                data: {
                    UserName: this.userName,
                    Password: this.Password,
                    Remember: this.Remember
                },
            }).then(function (response) {
                if (response.data.code === 200) {
                    location.href = response.data.obj;
                    toastr.success(response.data.msg);
                } else {
                    toastr.error(response.data.msg);
                }
            }).catch(function (error) {
                console.log(error);
            });
        }
    }
});