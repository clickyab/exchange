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

    $("#device-H-RTB").attr("value" , $(window).height() ) ;
    $("#device-W-RTB").attr("value" , $(window).width() ) ;


    // Initialize of Ip and user-agent value
    // GET IP of user from api
    var ip = "127.0.0.1" ;
    $.getJSON( "https://freegeoip.net/json/",
        function(data){
           ip = data.ip ;
            $("#device-IP-RTB").attr("value" , ip );
            $("#device-IP-REST").attr("value" ,  ip);
        }
    );
    // set user agent
    $("#useragent-RTB").attr("value" , navigator.userAgent );
    $("#device-ua-REST").attr("value" , navigator.userAgent );

    // Create Dynamic form with specific name
    // this element will update when you update your form in html
    // remove name to avoid conflict of new element with SerializeObject plugin
    function createForm() {

        var createElement = $(".banner-input").clone();
        var inputElem = createElement.find("input");

        // for(var i=0 ; i<inputElem.length ; i++ ){
        //   $(inputElem[i]).removeAttr("name");
        // }
        return '<div class="banner-input">' + createElement.html() + '<button type="button" class="btn-minus btn">- Remove Impression</button></div>';
    }


    function RandomIdGen(elem) {
        var elem = $(elem);
        for (var i = 0; i < elem.length; i++) {
            $(elem[i]).attr("value",(Math.ceil(Math.random() * 100000000000)));
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
                $((elem)[i]).attr("value" , array[i]) ;
            }
        }
        else {
            console.log("Error: Lenght of Element and Result not equal")
        }
    }


    function createJsonIMPORTB(elem){

       obj = {} ;
       obj.id = $(elem).find("[data-impId]").val();
       obj.banner= {} ;
       obj.banner.w         = $(elem).find("[data-impBannerW]").val();
       obj.banner.h         = $(elem).find("[data-impBannerH]").val();
       obj.banner.id        = $(elem).find("[data-impBannerId]").val();
       obj.banner.btype     = $(elem).find("[data-impBannerBtype]").val();
       obj.banner.battr     = $(elem).find("[data-impBannerBattr]").val();
       obj.banner.mimes     = $(elem).find("[data-impBannerMimes]").val();
       obj.bidfloor         = $(elem).find("[data-bidfloor]").val();
       obj.bidfloorcur      = $(elem).find("[data-bidfloorcur]").val();

       obj.banner.btype  = obj.banner.btype.split(",") ;
       obj.banner.battr  = obj.banner.battr.split(",") ;
       obj.banner.mimes  = obj.banner.mimes.split(",") ;

       ArrayElemToInt(obj.banner.btype);
       ArrayElemToInt(obj.banner.battr);



       if($(elem).find("[data-secure]:checked").length ===1){
           obj.secure = 1 ;
       }
       else{
           obj.secure = 0 ;
       }
       return obj ;
    }

    function createJsonIMPREST(elem){

        obj = {} ;
        obj.id = $(elem).find("[data-impId]").val();
        obj.banner= {} ;
        obj.banner.w         = $(elem).find("[data-impBannerW]").val();
        obj.banner.h         = $(elem).find("[data-impBannerH]").val();
        obj.banner.id        = $(elem).find("[data-impBannerId]").val();
        obj.bidfloor         = $(elem).find("[data-bidfloor]").val();
        obj.bidfloorcur      = $(elem).find("[data-bidfloorcur]").val();



        if($(elem).find("[data-secure]:checked").length ===1){
            obj.secure = 1 ;
        }
        else{
            obj.secure = 0 ;
        }
        return obj ;
    }

    function toIntAndEmpty(obj) {
        obj = parseInt(obj) || "";
    }

    function ArrayElemToInt(obj) {
        if (obj != "") {
            obj = obj.map(function (x) {
                return toIntAndEmpty(x);
            });
        }
    }

    // Get all impressions as object
    function getDynamicFormORTB(elem) {
        var obj = [];
        elemAppend = $(elem);
        for (var i = 0; i < elemAppend.length; i++) {
            obj[i] = (createJsonIMPORTB(elemAppend[i]));
        }
        return obj ;
    }
    function getDynamicFormREST(elem) {
        var obj = [];
        elemAppend = $(elem);
        for (var i = 0; i < elemAppend.length; i++) {
            obj[i] = (createJsonIMPREST(elemAppend[i]));
        }
        return obj ;
    }

    function BannerSplitor(elemSelector){
        SplitElems = $(elemSelector);
        for (i = 0 ; i< SplitElems.length ; i++){
            var impBanner =  $(".imp-banner") ;
            var splitedBanner = ($(SplitElems[i]).val()).split("x");
            impBanner[i] = splitedBanner;
            $(SplitElems[i]).parent().find("[data-impBannerW]").attr("value" , (impBanner[i])[0]) ;
            $(SplitElems[i]).parent().find("[data-impBannerH]").attr("value" , (impBanner[i])[1]) ;
        }
    }
    // append form action
    $(".banner-input-wrapper button").on("click", function (event) {
        appendElem = createForm();
        event.preventDefault();
        $(this).parent().append(appendElem);
    });
    // remove form action
    $("body").on("click", ".btn-minus", function (event) {
        event.preventDefault();
        $(this).parent().remove();
    });


    $(".RTB-form").submit(function (event) {
        event.preventDefault();
        BannerSplitor("#RTB-BannerSize") ;
        obj = getDynamicFormORTB(".RTB-form .banner-input");


        // Shortcut to add splited array of bcat input to object
        var jsonForm = $(this).serializeObject();
        var splitedBcat = splitElemVal("#bcat-RTB", ",");
        var splitedWlang = splitElemVal("#wlang-RTB", ",");
        jsonForm.bcat = splitedBcat;
        jsonForm.wlang = splitedWlang;
        jsonForm.imp = obj ;
        var impLength = obj.length;
        for (var i=0 ; i < impLength ; i++)
        {
            toIntAndEmpty(jsonForm.imp[i].banner.w);
            toIntAndEmpty(jsonForm.imp[i].banner.h);
            toIntAndEmpty(jsonForm.imp[i].bidfloor);
            toIntAndEmpty(jsonForm.imp[i].bidfloorcur);
        }

        jsonForm.device.mccmnc = (jsonForm.device.mccmnc).split("-");

        toIntAndEmpty(jsonForm.device.geo.lat);
        toIntAndEmpty(jsonForm.device.geo.lon );
        toIntAndEmpty(jsonForm.device.geo.type );
        toIntAndEmpty(jsonForm.device.geo.accuracy);


        toIntAndEmpty(jsonForm.device.h);
        toIntAndEmpty(jsonForm.device.w);

        toIntAndEmpty(jsonForm.tmax);
        toIntAndEmpty(jsonForm.test);
        toIntAndEmpty(jsonForm.at);


        jsonForm.wseat  = (jsonForm.wseat).split(",");
        jsonForm.bseat  = (jsonForm.bseat).split(",");
        jsonForm.cur    = (jsonForm.cur).split(",");
        jsonForm.badv   = (jsonForm.badv).split(",");
        jsonForm.bapp   = (jsonForm.bapp).split(",");

        delete jsonForm.radioSite;
        if(site){
           delete jsonForm.app;
           jsonForm.site.cat              = (jsonForm.site.cat).split(",");
           jsonForm.site.sectioncat       = (jsonForm.site.sectioncat).split(",");
           jsonForm.site.pagecat          = (jsonForm.site.pagecat).split(",");
           jsonForm.site.publisher.cat    = (jsonForm.site.publisher.cat).split(",");
        }
        else{
            delete jsonForm.site;
            jsonForm.app.cat              = (jsonForm.app.cat).split(",");
            jsonForm.app.sectioncat       = (jsonForm.app.sectioncat).split(",");
            jsonForm.app.pagecat          = (jsonForm.app.pagecat).split(",");
            jsonForm.app.publisher.cat    = (jsonForm.app.publisher.cat).split(",");
        }
        var JSONOut = {};
        JSONOut.request = jsonForm;
        JSONOut.meta = {};
        JSONOut.meta.key = $("#key-RTB").val() ;



        var url =$(".RTB-form").attr("action") ;

        $.ajax({
            url: url,
            type: 'POST',
            data: JSON.stringify(JSONOut),
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            async: false,
            success: function(msg) {
                alert(msg);
            }
        });

        console.log(JSON.stringify(JSONOut));

    });


    $(".REST-form").submit(function (event) {
        event.preventDefault();
        BannerSplitor("#REST-BannerSize") ;
        obj = getDynamicFormREST(".REST-form .banner-input");


        // Shortcut to add splited array of bcat input to object
        var jsonForm = $(this).serializeObject();
        var splitedBcat = splitElemVal("#bcat-REST", ",");
        jsonForm.bcat = splitedBcat;
        jsonForm.imp = obj ;
        var impLength = obj.length;
        for (var i=0 ; i < impLength ; i++)
        {
            toIntAndEmpty(jsonForm.imp[i].banner.w);
            toIntAndEmpty(jsonForm.imp[i].banner.h);
            toIntAndEmpty(jsonForm.imp[i].bidfloor);
            toIntAndEmpty(jsonForm.imp[i].bidfloorcur);

        }



        toIntAndEmpty(jsonForm.tmax);
        toIntAndEmpty(jsonForm.test);

        var JSONOut = {} ;
        JSONOut.request = jsonForm;
        JSONOut.meta = {} ;
        JSONOut.meta.key = $("#key-REST").val() ;
       var url =$(".REST-form").attr("action") ;


        $.ajax({
            url: url,
            type: 'POST',
            data: JSON.stringify(JSONOut),
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            async: false,
            success: function(msg) {
                alert(msg);
            }
        });

        console.log(JSON.stringify(JSONOut));

    });
});