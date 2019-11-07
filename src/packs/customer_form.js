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
            Id: '',
            Address: '',
            Website: '',
            Amount: 0,
            Aging: 0,
            CompanyType: '',
            IsVip: false,
            SaleUser: '',
            AuditUser: '',

        },
        select2Data: [],
        options: [],
        userOptions: [],
        selectCompanyType: {},
        selectPeriod: {},
        selectVip: {},

    },
    mounted: function () {
        this.initData();
    },
    methods: {
        select2Method: function (actionType) {
            let data = window.selectApi("/customer/get_status", {actionType: actionType}, 1);
            this.select2Data = data

        },
        clickCompany: function () {
            this.options = this.$select2Company();
        },
        clickUser: function () {
            let str = "";
            let company = this.customer.Company;
            if (company !== '' && company !== null) {
                str = "Company:" + company.Id;
            }
            this.userOptions = this.$select2User({query: str})
        },
        initData: function () {
            let _this = this;
            let location_url = location.pathname.split('/');
            let id = location_url[location_url.length - 1];
            console.log(id);
            if (id !== "" && typeof (id) !== "undefined") {
                this.$showInitData("/customer/" + id).then(res => {
                        console.log(res);
                        _this.customer = res.data;
                        _this.selectVip = res['other']['IsVip'];
                        _this.selectCompanyType = res['other']['selectCompanyType'];
                        _this.selectPeriod = res['other']['selectPeriod'];
                        console.log(_this.selectPeriod);
                    },
                    error => {
                        console.log(error);
                    });
            }
        },
        submitForm: function () {
            let _this = this;
            let url, method;
            if (this.customer.Id === '') {
                url = "/customer";
                method = 'post';
            } else {
                url = "/customer/" + this.customer.Id;
                method = 'put';
            }
            if (this.validateForm()) {
                this.defaultValue();
                url = url + "?audit_id=" + this.customer.AuditUser.Id;
                url = url + "&sale_id=" + this.customer.SaleUser.Id;
                url = url + "&period=" + this.customer.AccountPeriod;
                url = url + '&company_id=' + this.customer.Company.Id;
                url = url + "&company_type=" + this.customer.CompanyType;
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
            if (!this.selectPeriod) {
                this.errors.push("账期不能为空");
                return false;
            }
            if (!this.customer.Company) {
                this.errors.push("所属公司不能为空");
                return false;
            }
            if (!this.selectCompanyType) {
                this.errors.push("类型不能为空");
                return false;
            }
            if (!this.customer.SaleUser) {
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
            if (this.selectPeriod) {
                this.customer.AccountPeriod = this.selectPeriod.code;
            }
            if (this.selectCompanyType) {
                this.customer.CompanyType = this.selectCompanyType.code;
            }
            if (this.selectVip) {
                this.customer.IsVip = this.selectVip.code;
            }
            this.customer.Aging = parseInt(this.customer.Aging);
            this.customer.Amount = parseInt(this.customer.Amount);
        }
    }
});