$(function () {



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


    function createJsonObj(elem){

        obj = {} ;
        obj.id = $(elem).find("[data-impId]").val();
        obj.banner= {} ;
        obj.banner.w      = $(elem).find("[data-impBannerW]").val();
        obj.banner.h      = $(elem).find("[data-impBannerH]").val();
        obj.banner.id     = $(elem).find("[data-impBannerId]").val();
        obj.bidfloor      = $(elem).find("[data-bidfloor]").val();
        obj.bidfloorcur   = $(elem).find("[data-bidfloorcur]").val();

        if($(elem).find("[data-secure]:checked").length ===1){
            obj.secure = 1 ;
        }
        else{
            obj.secure = 0 ;
        }
        return obj ;
    }

    // Get all impressions as object
    function getDynamicForm(elem) {
        var obj = [];
        elemAppend = $(elem);
        for (var i = 0; i < elemAppend.length; i++) {
            obj[i] = (createJsonObj(elemAppend[i]));
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
        console.log($(this).parent());
        $(this).parent().remove();
    });


    $(".RTB-form").submit(function (event) {
        event.preventDefault();
        RandomIdGen(".randomField");
        BannerSplitor("#RTB-BannerSize") ;
        obj = getDynamicForm(".RTB-form .banner-input");


        // Shortcut to add splited array of bcat input to object
        var jsonForm = $(this).serializeObject();
        var splitedBcat = splitElemVal("#bcat", ",");
        var splitedWlang = splitElemVal("#wlang", ",");
        jsonForm.request.bcat = splitedBcat;
        jsonForm.request.wlang = splitedWlang;
        jsonForm.request.imp = obj ;
        var impLength = obj.length;
        for (var i=0 ; i < impLength ; i++)
        {
            jsonForm.request.imp[i].banner.w = parseInt( jsonForm.request.imp[i].banner.w);
            jsonForm.request.imp[i].banner.h = parseInt( jsonForm.request.imp[i].banner.h);
            jsonForm.request.imp[i].bidfloor = parseInt( jsonForm.request.imp[i].bidfloor);
            jsonForm.request.imp[i].bidfloorcur = parseInt( jsonForm.request.imp[i].bidfloorcur);
        }


        jsonForm.request.tmax = parseInt( jsonForm.tmax);
        jsonForm.request.test= parseInt( jsonForm.test);

        var url = $(".RTB-form").attr("action");

        $.ajax({
            url: url,
            type: 'POST',
            data: JSON.stringify(jsonForm),
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            async: false,
            success: function(msg) {
                alert(msg);
            }
        });

        console.log(JSON.stringify(jsonForm));

    });




    $(".REST-form").submit(function (event) {
        event.preventDefault();
        RandomIdGen(".randomField");
        BannerSplitor("#REST-BannerSize") ;
        obj = getDynamicForm(".REST-form .banner-input");


        // Shortcut to add splited array of bcat input to object
        var jsonForm = $(this).serializeObject();
        var splitedBcat = splitElemVal("#bcat-REST", ",");
        var splitedBadv = splitElemVal("#badv-REST", ",");
        var splitedBseat = splitElemVal("#bseat-REST", ",");
        jsonForm.bcat = splitedBcat;
        jsonForm.badv = splitedBadv;
        jsonForm.bseat = splitedBseat;
        jsonForm.imp = obj ;
        var impLength = obj.length;
        for (var i=0 ; i < impLength ; i++)
        {
            jsonForm.imp[i].banner.w = parseInt( jsonForm.imp[i].banner.w);
            jsonForm.imp[i].banner.h = parseInt( jsonForm.imp[i].banner.h);
            jsonForm.imp[i].bidfloor = parseInt( jsonForm.imp[i].bidfloor);
            jsonForm.imp[i].bidfloorcur = parseInt( jsonForm.imp[i].bidfloorcur);
        }

        jsonForm.tmax = parseInt( jsonForm.tmax);
        jsonForm.test= parseInt( jsonForm.test);

        var url = $(".REST-form").attr("action");

        $.ajax({
            url: url,
            type: 'POST',
            data: JSON.stringify(jsonForm),
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            async: false,
            success: function(msg) {
                alert(msg);
            }
        });

        console.log(JSON.stringify(jsonForm));

    });
});