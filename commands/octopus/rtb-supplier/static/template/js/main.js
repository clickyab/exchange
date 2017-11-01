$(function () {

    RandomIdGen(".randomField");


    // Show SITE OR APP FOR VIA RADIO BTN CHANGE
    var site = true;
    SelectAppSiteFormRTB();
    SelectAppSiteFormREST();


    $(document).on("change", ".site-app-RTB", function () {
        SelectAppSiteFormRTB();
    });
    $(document).on("change", ".site-app-REST", function () {
        SelectAppSiteFormREST();
    });

    function SelectAppSiteFormRTB() {
        $("[data-SiteSelect]:checked").each(function () {
                $(".site-form-RTB").show();
                $(".app-form-RTB").hide();
                site = true;
            }
        );
        $("[data-AppSelect]:checked").each(function () {
            $(".site-form-RTB").hide();
            $(".app-form-RTB").show();
            site = false;
        });
    }

    function SelectAppSiteFormREST() {
        $("[data-SiteSelect-REST]:checked").each(function () {
                $(".site-form-REST").show();
                $(".app-form-REST").hide();
                site = true;
            }
        );
        $("[data-AppSelect-REST]:checked").each(function () {
            $(".site-form-REST").hide();
            $(".app-form-REST").show();
            site = false;
        });
    }


    // Get Width and Height of Device

    $("#device-H-RTB").attr("value", $(window).height());
    $("#device-W-RTB").attr("value", $(window).width());


    // Initialize of Ip and user-agent value
    // GET IP of user from api
    var ip = "127.0.0.1";
    $.getJSON("https://freegeoip.net/json/",
        function (data) {
            ip = data.ip;
            $("#device-IP-RTB").attr("value", ip);
            $("#device-IP-REST").attr("value", ip);
        }
    );
    // set user agent
    $("#useragent-RTB").attr("value", navigator.userAgent);
    $("#device-ua-REST").attr("value", navigator.userAgent);

    // Create Dynamic form with specific name
    // this element will update when you update your form in html
    // remove name to avoid conflict of new element with SerializeObject plugin
    function createFormORTB() {

        var createElement = $(".RTB-banner-input").clone();
        var inputElem = createElement.find("input");

        return '<div class="RTB-banner-input">' + createElement.html() + '<button type="button" class="btn-minus btn">- Remove Impression</button></div>';
    }

    function createFormREST() {

        var createElement = $(".REST-banner-input").clone();
        var inputElem = createElement.find("input");

        return '<div class="REST-banner-input">' + createElement.html() + '<button type="button" class="btn-minus btn">- Remove Impression</button></div>';
    }


    function RandomIdGen(elem) {
        var elem = $(elem);
        for (var i = 0; i < elem.length; i++) {
                $(elem[i]).attr("value", (Math.ceil(Math.random() * 100000000000)));
        }
    }

    function splitElemVal(elem, splitor) {
        var value = $(elem).val();
        var result = value.split(splitor);
        return result;
    }

    function fillValue(elem, array) {
        if (elem.length === array.length) {
            for (var i = 0; i < elem.length; i++) {
                $((elem)[i]).attr("value", array[i]);
            }
        }
        else {
            console.log("Error: Lenght of Element and Result not equal")
        }
    }


    function createJsonIMPORTB(elem) {

        obj = {};
        obj.id = $(elem).find("[data-impId]").val();
        obj.banner = {};
        obj.banner.w = $(elem).find("[data-impBannerW]").val();
        obj.banner.h = $(elem).find("[data-impBannerH]").val();
        obj.banner.id = $(elem).find("[data-impBannerId]").val();
        obj.banner.btype = $(elem).find("[data-impBannerBtype]").val();
        obj.banner.battr = $(elem).find("[data-impBannerBattr]").val();
        obj.banner.mimes = $(elem).find("[data-impBannerMimes]").val();
        obj.bidfloor = $(elem).find("[data-bidfloor]").val();
        obj.bidfloorcur = $(elem).find("[data-bidfloorcur]").val();


        obj.banner.w = parseInt(obj.banner.w);
        obj.banner.h = parseInt(obj.banner.h);


        if (obj.banner.btype != "") {
            obj.banner.btype = obj.banner.btype.split(",");
            obj.banner.btype = obj.banner.btype.map(function (x) {
                return parseInt(x);
            });
        }
        if (obj.banner.battr != "") {
            obj.banner.battr = obj.banner.battr.split(",");
            obj.banner.battr = obj.banner.battr.map(function (x) {
                return parseInt(x);
            });
        }

        if (obj.banner.mimes != "") {
            obj.banner.mimes = obj.banner.mimes.split(",");
        }


        if ($(elem).find("[data-secure]:checked").length === 1) {
            obj.secure = 1;
        }
        else {
            obj.secure = 0;
        }
        return obj;
    }

    function createJsonIMPREST(elem) {

        obj = {};
        obj.id = $(elem).find("[data-impId]").val();
        obj.banner = {};
        obj.banner.w = $(elem).find("[data-impBannerW]").val();
        obj.banner.h = $(elem).find("[data-impBannerH]").val();
        obj.banner.id = $(elem).find("[data-impBannerId]").val();
        obj.bidfloor = $(elem).find("[data-bidfloor]").val();

        obj.banner.w = parseInt(obj.banner.w);
        obj.banner.h = parseInt(obj.banner.h);


        if ($(elem).find("[data-secure]:checked").length === 1) {
            obj.secure = 1;
        }
        else {
            obj.secure = 0;
        }
        return obj;
    }


    // Get all impressions as object
    function getDynamicFormORTB(elem) {
        var obj = [];
        elemAppend = $(elem);
        console.log(elemAppend.length);
        for (var i = 0; i < elemAppend.length; i++) {
            obj[i] = (createJsonIMPORTB(elemAppend[i]));
        }
        return obj;
    }

    function getDynamicFormREST(elem) {
        var obj = [];
        elemAppend = $(elem);
        for (var i = 0; i < elemAppend.length; i++) {
            obj[i] = (createJsonIMPREST(elemAppend[i]));
        }
        return obj;
    }

    function BannerSplitor(elemSelector) {
        SplitElems = $(elemSelector);
        for (i = 0; i < SplitElems.length; i++) {
            var impBanner = $(".imp-banner");
            var splitedBanner = ($(SplitElems[i]).val()).split("x");
            impBanner[i] = splitedBanner;
            $(SplitElems[i]).parent().find("[data-impBannerW]").attr("value", (impBanner[i])[0]);
            $(SplitElems[i]).parent().find("[data-impBannerH]").attr("value", (impBanner[i])[1]);
        }
    }

    function clean(obj) {
        var propNames = Object.getOwnPropertyNames(obj);
        for (var i = 0; i < propNames.length; i++) {
            var propName = propNames[i];
            if (obj[propName] === null  || obj[propName] === undefined || obj[propName] === "") {
                delete obj[propName];
            }
        }
    }


    // append form action
    $(".RTB-banner-input-wrapper button").on("click", function (event) {
        appendElem = createFormORTB();
        event.preventDefault();
        $(this).parent().append(appendElem);
        RandomIdGen(".randomField");
    });
    $(".REST-banner-input-wrapper button").on("click", function (event) {
        appendElem = createFormREST();
        event.preventDefault();
        $(this).parent().append(appendElem);
        RandomIdGen(".randomField");
    });


    // remove form action
    $("body").on("click", ".btn-minus", function (event) {
        event.preventDefault();
        $(this).parent().remove();
    });


    $(".RTB-form").submit(function (event) {
        event.preventDefault();
        BannerSplitor(".RTB-BannerSize");
        obj = getDynamicFormORTB(".RTB-form .RTB-banner-input");


        // Shortcut to add splited array of bcat input to object
        var jsonForm = $(this).serializeObject();
        jsonForm.imp = obj;
        var impLength = obj.length;
        for (var i = 0; i < impLength; i++) {
            jsonForm.imp[i].banner.w = parseInt(jsonForm.imp[i].banner.w);
            jsonForm.imp[i].banner.h = parseInt(jsonForm.imp[i].banner.h);
            jsonForm.imp[i].bidfloor = parseInt(jsonForm.imp[i].bidfloor) || "";
            jsonForm.imp[i].bidfloorcur = parseInt(jsonForm.imp[i].bidfloorcur) || "";
            clean(jsonForm.imp[i]);
            clean(jsonForm.imp[i].banner);

        }

        if (jsonForm.bcat !== "") {
            jsonForm.bcat = (jsonForm.bcat).split(",");
        }
        if (jsonForm.wlang !== "") {
            jsonForm.wlang = (jsonForm.wlang).split(",");
        }
        jsonForm.device.mccmnc = (jsonForm.device.mccmnc).split("-");

        jsonForm.device.geo.lat = parseInt(jsonForm.device.geo.lat) || "";
        jsonForm.device.geo.lon = parseInt(jsonForm.device.geo.lon) || "";
        jsonForm.device.geo.type = parseInt(jsonForm.device.geo.type) || "";
        jsonForm.device.geo.accuracy = parseInt(jsonForm.device.geo.accuracy) || "";


        jsonForm.device.h = parseInt(jsonForm.device.h);
        jsonForm.device.w = parseInt(jsonForm.device.w);
        jsonForm.device.connectiontype = parseInt(jsonForm.device.connectiontype);

        jsonForm.tmax = parseInt(jsonForm.tmax);
        jsonForm.test = parseInt(jsonForm.test);
        jsonForm.at = parseInt(jsonForm.at);


        if (jsonForm.wseat !== "") {
            jsonForm.wseat = (jsonForm.wseat).split(",");
        }
        if (jsonForm.bseat !== "") {
            jsonForm.bseat = (jsonForm.bseat).split(",");
        }
        if (jsonForm.cur !== "") {
            jsonForm.cur = (jsonForm.cur).split(",");
        }
        if (jsonForm.badv !== "") {
            jsonForm.badv = (jsonForm.badv).split(",");
        }
        if (jsonForm.bapp !== "") {
            jsonForm.bapp = (jsonForm.bapp).split(",");
        }

        delete jsonForm.radioSite;
        if (site) {
            delete jsonForm.app;
            if (jsonForm.site.cat !== "") {
                jsonForm.site.cat = (jsonForm.site.cat).split(",");
            }
            if (jsonForm.site.sectioncat !== "") {
                jsonForm.site.sectioncat = (jsonForm.site.sectioncat).split(",");
            }
            if (jsonForm.site.pagecat !== "") {
                jsonForm.site.pagecat = (jsonForm.site.pagecat).split(",");
            }
            if (jsonForm.site.publisher.cat !== "") {
                jsonForm.site.publisher.cat = (jsonForm.site.publisher.cat).split(",");
            }
            clean((jsonForm.site.publisher));
            clean((jsonForm.site));

        }
        else {
            delete jsonForm.site;
            if (jsonForm.app.cat !== "") {
                jsonForm.app.cat = (jsonForm.app.cat).split(",");
            }
            if (jsonForm.app.sectioncat !== "") {
                jsonForm.app.sectioncat = (jsonForm.app.sectioncat).split(",");
            }
            if (jsonForm.app.pagecat !== "") {
                jsonForm.app.pagecat = (jsonForm.app.pagecat).split(",");
            }
            if (jsonForm.app.publisher.cat !== "") {
                jsonForm.app.publisher.cat = (jsonForm.app.publisher.cat).split(",");
            }
            clean((jsonForm.app.publisher));
            clean((jsonForm.app));
        }
        clean(jsonForm);
        clean((jsonForm.device.geo));
        clean((jsonForm.device.ext));
        clean((jsonForm.device));

        var JSONOut = {};
        JSONOut.request = jsonForm;
        JSONOut.meta = {};
        JSONOut.meta.key = $("#key-RTB").val();


        var url = $(".RTB-form").attr("action");
        var jsonRes;

        $.ajax({
            url: url,
            type: 'POST',
            data: JSON.stringify(JSONOut),
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            async: false,
            success: function (msg) {
                jsonRes = msg;
                $(".json-out-RTB").html( syntaxHighlight(JSON.stringify(jsonRes, undefined, 4)));
            }
        });

        console.log(JSON.stringify(JSONOut));



    });


    $(".REST-form").submit(function (event) {
        event.preventDefault();
        BannerSplitor(".REST-BannerSize");
        obj = getDynamicFormREST(".REST-form .REST-banner-input");


        var jsonForm = $(this).serializeObject();
        jsonForm.imp = obj;
        var impLength = obj.length;
        for (var i = 0; i < impLength; i++) {

            jsonForm.imp[i].banner.w = parseInt(jsonForm.imp[i].banner.w);
            jsonForm.imp[i].banner.h = parseInt(jsonForm.imp[i].banner.h);
            jsonForm.imp[i].bidfloor = parseInt(jsonForm.imp[i].bidfloor) || "";
            clean(jsonForm.imp[i]);
        }


        if (jsonForm.bcat !== "") {
            jsonForm.bcat = (jsonForm.bcat).split(",");
        }
        jsonForm.device.h = parseInt(jsonForm.device.h) || "";
        jsonForm.device.w = parseInt(jsonForm.device.w) || "";
        jsonForm.device.connectiontype = parseInt(jsonForm.device.connectiontype);

        jsonForm.tmax = parseInt(jsonForm.tmax) || "";
        jsonForm.test = parseInt(jsonForm.test) || "";
        jsonForm.at = parseInt(jsonForm.at) || "";

        delete jsonForm.radioSite;
        if (site) {
            delete jsonForm.app;
            if (jsonForm.site.cat !== "") {
                jsonForm.site.cat = (jsonForm.site.cat).split(",");
            }
            clean((jsonForm.site));
        }
        else {
            delete jsonForm.site;
            if (jsonForm.app.cat !== "") {
                jsonForm.app.cat = (jsonForm.app.cat).split(",");
            }
            clean((jsonForm.app));

        }
        clean(jsonForm);
        clean((jsonForm.device));

        var JSONOut = {};
        JSONOut.request = jsonForm;
        JSONOut.meta = {};
        JSONOut.meta.key = $("#key-REST").val();
        var url = $(".REST-form").attr("action");
        var jsonRes;


        $.ajax({
            url: url,
            type: 'POST',
            data: JSON.stringify(JSONOut),
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            async: false,
            success: function (msg) {
                jsonRes = msg;
                $(".json-out-REST").html(syntaxHighlight(JSON.stringify(jsonRes, undefined, 4)));

            }
        });

        console.log(JSON.stringify(JSONOut));

    });


    function syntaxHighlight(json) {
        json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
        return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
            var cls = 'number';
            if (/^"/.test(match)) {
                if (/:$/.test(match)) {
                    cls = 'key';
                } else {
                    cls = 'string';
                }
            } else if (/true|false/.test(match)) {
                cls = 'boolean';
            } else if (/null/.test(match)) {
                cls = 'null';
            }
            return '<span class="' + cls + '">' + match + '</span>';
        });
    }

});