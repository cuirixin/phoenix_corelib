/*
* @Author: rixincui
* @Email: 277831894@qq.com
* @Date:  2016-10-18 16:30:51
* @Last Modified by: 
* @Last Modified time: 2016-10-18 16:37:23
* @Descrition : 阿里云 api 网关签名
*/
"use strict";
var querystring = require('querystring');
var crypto = require('crypto');
var url = require('url');


module.exports = function(config, requestParams, callBack ){

    var config = config || {
        ali_gateway_key: "",
        ali_gateway_secret: ""
    }

    // 定义 key 和 secret
    var KEY = config.ali_gateway_key;
    var SECRET = config.ali_gateway_secret;

    // Content-Type、Accept、Content-MD5、Date, 这三个是基础的签名字符串， 必须包含
    var defaultSignObj = {
        'Content-Type' : 'application/x-www-form-urlencoded;charset=UTF-8',
        'Accept' : 'application/json; charset=UTF-8',
        'Content-MD5' : '',
        'Date' : new Date(),
        'Method' : 'GET',
        'X-Ca-Stage' : config.ali_stage || 'TEST',
        'X-Ca-Key': KEY,
        'X-Ca-Nonce' : getNonce(),
        'X-Ca-Timestamp' : new Date().getTime(),
        //'V-App-Client-Information' : getInfomation(), //此字段根据自己的 api 的定义选择传递
        'X-Ca-Signature-Headers' : 'X-Ca-Key,X-Ca-Nonce,X-Ca-Stage,X-Ca-Timestamp',
    };

    for (var k in requestParams) {
        defaultSignObj[k] = requestParams[k];
    };

    requestParams = defaultSignObj;
    var stringToSign = requestParams.Method.toUpperCase()+"\n"+requestParams.Accept+"\n"+requestParams['Content-MD5']+"\n"+requestParams['Content-Type']+"\n"+requestParams.Date+"\n";

    // 检查参数签名的定义参数, 获取签名的参数
    var signatureHeaders = requestParams['X-Ca-Signature-Headers'].split(',');
    if(signatureHeaders.length > 0){
        var Headers = {};
        signatureHeaders.forEach(function(val, key){
            Headers[val] = requestParams[val];
        });
        // 按照字典对 Key 进行排序
        var keys = Object.keys(Headers);
        var newKeys = keys.sort();
        var newHeaders = {};
        newKeys.forEach(function(val, key){
            stringToSign += val + ':' +requestParams[val]+"\n";
        });

        requestParams = getQueryOfUrl(requestParams);

        stringToSign += requestParams.Path;
        // 获取签名
        requestParams['X-Ca-Signature'] = getSignature(SECRET, stringToSign);
        
        //  签名后， 发送前，处理中文编码。
        requestParams.Path = encodeURI(requestParams.Path);
        requestParams.Url = encodeURI(requestParams.Url);


        var finalRequestParams = {
            method: requestParams.Method,
            url: requestParams.Url,
            form: requestParams.Form,  
        }

        delete requestParams.Form;
        finalRequestParams['headers'] = requestParams;


        callBack(null, finalRequestParams);
    } else {
        callBack('X-Ca-Signature-Headers is required');
    }
};

/**
 * 获取请求头部的 V-App-Client-Information 的信息
 * @return {[type]} [description]
 */
function getInfomation(){
    return 'app_name:hxwx|plat:win32|ver:3.3|device:wap|os:node|channel_name:wap|udid:"+new Date().getTime()+"|client_ip:192.168.0.1|user-agent:test';
}

/**
 * 获取 url的query 部分
 *
 * @param  {[type]} requestParams [description]
 * @return {[type]}               [description]
 */
function getQueryOfUrl (requestParams){
    // 如果有参数传递， 测排序后拼接参数
    if( requestParams.Form ){
        var newForm = {};
        requestParams.Url += '?';
        // 按照字典对 Key 进行排序
        Object.keys(requestParams.Form).sort().forEach(function(bVal, bKey){
             newForm[bVal] = requestParams.Form[bVal];
        });

        for (var key in newForm){
            if (newForm[key]) {
                requestParams.Url += key + '=' + newForm[key] + '&';
            } else {
                requestParams.Url += key  + '&';
            }
        }
        // 删除最后一位的 &
        requestParams.Url = requestParams.Url.substring(0, requestParams.Url.length -1);
        // requestParams.Url += '?' + querystring.unescape(querystring.stringify(newForm));
    }
    // 此处 url.parse 函数会把原有的字符串编码，所以此处需要减编码后的字符串复原。
    requestParams.Path = decodeURI(url.parse(requestParams.Url).path);

    return requestParams;
}

/**
 * 请求唯一标识，15分钟内AppKey+API+Nonce不能重复，与时间戳结合使用才能起到防重放作用
 * @return {[type]} [description]
 */
function getNonce(){
    return Math.random().toString(36).substr(2, 15);
    // 此处如果 nonce 硬编码了， 接口会返回 Nonce used。
}

/**
 * 获取签名
 * @param  {[type]} stringToSign [description]
 * @return {[type]}              [description]
 */
function getSignature(SECRET, stringToSign){
    return crypto.createHmac('sha256',SECRET).update(stringToSign, 'utf8').digest('base64');
}