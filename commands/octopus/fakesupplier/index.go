package main

const index = `<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <title>Document</title>
    <style>
        input.inner-input{
            padding: 10px;
        }
        *{
            font-size: 12px ;
        }
    </style>
</head>
<body>
    <div class="container" style="margin-top: 20px;">



        <form method="post" action="send" class="form-horizontal final-form">
            <div class="row">
                <div class="form-group">
                    <label for="type">Type</label>
                    <input type="text" class="form-control" id="type" name="type">
                </div>
                <div class="form-group">
                    <label for="scheme">Scheme</label>
                    <input type="text" class="form-control" id="scheme" name="scheme">
                </div>
                <div class="form-group">
                    <label for="page_track_id">PageTrackID</label>
                    <input type="text" class="form-control" id="page_track_id" name="page_track_id">
                </div>
                <div class="form-group">
                    <label for="user_track_id">UserTrackID</label>
                    <input type="text" class="form-control" id="user_track_id" name="user_track_id">
                </div>
                <div class="form-group">
                    <label for="refferer">Refferer</label>
                    <input type="text" class="form-control" id="refferer" name="refferer">
                </div>
                <div class="form-group">
                    <label for="parent">Parent</label>
                    <input type="text" class="form-control" id="parent" name="parent">
                </div>
                <div class="form-group">
                    <label for="user_agent">UserAgent</label>
                    <input type="text" class="form-control" id="user_agent" name="user_agent">
                </div>
                <div class="form-group">
                    <label for="ip">IP</label>
                    <input type="text" class="form-control" id="ip" name="ip">
                </div>
                <div class="form-group">
                    <label for="soft_floor">Soft Floor CPM</label>
                    <input type="number" class="form-control" id="soft_floor" name="soft_floor">
                </div>
                <div class="form-group">
                    <label for="floor_cpm">Floor CPM</label>
                    <input type="number" class="form-control" id="floor_cpm" name="floor_cpm">
                </div>
                <div class="form-group">
                    <label for="publisher_name">Publisher Name</label>
                    <input type="text" class="form-control" id="publisher_name" name="publisher_name">
                </div>
                <div class="form-group">
                    <label for="categories">Categories (comma, separated)</label>
                    <input type="text" class="form-control" id="categories" name="categories">
                </div>
                <div class="col-md-2">
                    <div class="panel panel-default my-panel">
                        <input class="track" type="hidden" name="track[]">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width1">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width1" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height1">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height1" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width2">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width2" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height2">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height2" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width3">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width3" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height3">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height3" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width4">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width4" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height4">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height4" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width5">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width5" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height5">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height5" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width6">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width6" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height6">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height6" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width7">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width7" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height7">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height7" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width8">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width8" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height8">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height8" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width9">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width9" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height9">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height9" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width10">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width10" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height10">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height10" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width11">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width11" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height11">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height11" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="col-md-2">
                    <input class="track" type="hidden" name="track[]">
                    <div class="panel panel-default my-panel">
                        <div class="panel-heading">
                            <h3 class="panel-title"> Slot Property</h3>
                        </div>
                        <div class="panel-body">
                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="width12">Width</label>
                                <br>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="width12" name="width[]">
                            </div>

                            <div style="padding: 10px;" class="form-group">
                                <label class="" for="height12">Height</label>
                                <input style="width: 140px;height: 20px;" type="number" class="inner-input" id="height12" name="height[]">
                            </div>

                        </div>
                    </div>
                </div>
                <div class="form-group">
                    <button type="submit" class="btn btn-default form-control">Apply</button>
                </div>
            </div>
        </form>
    </div>

    <script>
        function sha1()
        {
            var text = "";
            var possible = "abcdef0123456789";

            for( var i=0; i < 32; i++ )
                text += possible.charAt(Math.floor(Math.random() * possible.length));

            return text;
        }
        $(document).ready(function () {

            $(".track").each(function (element, v) {
                console.log(v)
                $(v).val(sha1())
            })

        })

    </script>
</body>
</html>
`