import{i as _}from "./8445cc3c.js";import{f as g,H as v,ad as b,A as F,l as B,j as l,u as e,o as D,i as r,ae as s,af as n,aq as N,e as C,ai as k,n as y,aj as A,ak as x}from "./315a852e.js";import{M as p}from "./index.148f4d09.js";const q={class:"flex justify-end"},w=g({__name:"Add",setup(E){const i=v(null),u=b({name:"",agent_id:0,corp_id:"",corp_secret:""}),m={name:{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0",trigger:"blur"},agent_id:{type:"number",required:!0,message:"\u8BF7\u8F93\u5165\u5E94\u7528ID",trigger:"blur"},corp_id:{required:!0,message:"\u8BF7\u8F93\u5165\u5E94\u7528Key",trigger:"blur"},corp_secret:{required:!0,message:"\u8BF7\u8F93\u5165\u5E94\u7528Secret",trigger:"blur"}},c=F();function f(){var o;(o=i.value)==null||o.validate(a=>{var t;a?(console.log("errors"),(t=p)==null||t.error("\u9A8C\u8BC1\u5931\u8D25")):_(u).then(I=>{var d;(d=p)==null||d.success("success"),c.push({name:"CredentialDingtalk"})})})}return(o, a)=>(D(),B(e(x),{title:"\u6DFB\u52A0DingTalk\u51ED\u8BC1"},{default:l(()=>[r(e(A),{ref_key:"formRef",ref:i,model:u,"label-placement":"left","label-width":160,rules:m,style:{maxWidth:"640px"}},{default:l(()=>[r(e(s),{label:"\u540D\u79F0:",path:"name"},{default:l(()=>[r(e(n),{value:u.name,"onUpdate:value":a[0]||(a[0]= t=>u.name=t)},null,8,["value"])]),_:1}),r(e(s),{label:"AgentId:",path:"agent_id"},{default:l(()=>[r(e(N),{value:u.agent_id,"onUpdate:value":a[1]||(a[1]= t=>u.agent_id=t),"show-button":!1},null,8,["value"])]),_:1}),r(e(s),{label:"AppKey:",path:"corp_id"},{default:l(()=>[r(e(n),{value:u.corp_id,"onUpdate:value":a[2]||(a[2]= t=>u.corp_id=t)},null,8,["value"])]),_:1}),r(e(s),{label:"AppSecret:",path:"corp_secret"},{default:l(()=>[r(e(n),{value:u.corp_secret,"onUpdate:value":a[3]||(a[3]= t=>u.corp_secret=t)},null,8,["value"])]),_:1}),C("div",q,[r(e(k),{onClick:f,type:"primary"},{default:l(()=>[y("\u63D0\u4EA4")]),_:1})])]),_:1},8,["model"])]),_:1}))}});export{w as default};