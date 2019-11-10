var app = new Vue({
    el: '#create_form',
    delimiters: ['{', '}'],
    data: {
        errors: [],
        customer: {
            Name: '',
            Telephone: '',
            Email: '',
            BusinessTypeName: '',
            Company: '',
            AccountPeriod: '',
            Address: '',
            Website: '',
            Amount: 0,
            Aging: 0,
            CompanyType: '',
            IsVip: false,
            SaleUser: '',
            AuditUser: '',

        },
        AccountPeriodOptions: [],
        CompanyOptions: [],
        userOptions: [],
        CompanyTypeOptions: [],
        IsVipOptions: [],
        Id: '',
        Url: '',
    },
    mounted: function () {
        this.confirmUrl();
        this.initData();
        this.select2Method();
    },
    methods: {
        confirmUrl: function () {
            if (location.pathname.indexOf("supplier") > -1) {
                this.Url = "/supplier"
            } else {
                this.Url = "/customer"
            }
        },
        select2Method: function () {
            let data = window.selectApi("/customer/status", {actionType: 'all'}, 1);
            this.AccountPeriodOptions = data.AccountPeriod;
            this.CompanyTypeOptions = data.CompanyType;
            this.IsVipOptions = data.IsVip;
            if (this.Url === '/supplier') {
                this.CompanyTypeOptions.splice(0, 1)
            } else {
                this.CompanyTypeOptions.splice(1, 1)
            }

        },
        clickCompany: function (search) {
            let options = {};
            if (search) {
                options['query'] = "Name:" + search;
            }
            this.CompanyOptions = this.$select2Company(options);
        },
        clickUser: function (search) {
            let str = "";
            let company = this.customer.Company;
            if (company !== '' && company !== null) {
                str = "Company:" + company.Id;
            }
            if (search) {
                str += "Name:" + search;
            }
            this.userOptions = this.$select2User({query: str})
        },
        initData: function () {
            let _this = this;
            let location_url = location.pathname.split('/');
            let id = location_url[location_url.length - 1];
            if (!isNaN(parseInt(id))) {
                this.$showInitData((this.Url + "/" + id)).then(res => {
                        console.log(res);
                        _this.customer = res;
                        _this.Id = res.Id;
                    },
                    error => {
                        console.log(error);
                    });
            }
        },
        submitForm: function () {
            let _this = this;
            let url, method;
            if (this.Id === '') {
                url = this.Url;
                method = 'post';
            } else {
                url = this.Url + "/" + this.Id;
                method = 'put';
            }
            if (this.validateForm()) {
                this.defaultValue();
                this.$submitFormData(method, url, this.customer).then(res => {
                        if (res.length !== 0) {
                            $.each(res, function (key, value) {
                                console.log(value);
                                _this.errors.push(key + ':' + value);
                            });
                        }
                    },
                    error => {
                        console.log(error);
                    });
            }
        },
        //验证表单是否通过
        validateForm: function () {
            this.errors = [];
            if (!this.customer.Name) {
                this.errors.push("客户名称不能为空");
                return false;
            }
            if (!this.customer.Telephone) {
                this.errors.push('联系电话不能为空');
                return false;
            }
            if (!this.customer.Email) {
                this.errors.push('邮箱不能为空');
                return false;
            }
            if (!this.customer.AccountPeriod) {
                this.errors.push("账期不能为空");
                return false;
            }
            if (!this.customer.Company) {
                this.errors.push("所属公司不能为空");
                return false;
            }
            if (!this.customer.CompanyType) {
                this.errors.push("类型不能为空");
                return false;
            }
            if (!this.customer.SaleUser && this.Url === "/customer") {
                this.errors.push("业务员不能为空");
                return false;
            }
            if (!this.customer.AuditUser) {
                this.errors.push("审核者不能为空");
                return false;
            }
            return true;
        },
        defaultValue: function () {
            this.customer.Aging = parseInt(this.customer.Aging);
            this.customer.Amount = parseInt(this.customer.Amount);
        }
    }
});