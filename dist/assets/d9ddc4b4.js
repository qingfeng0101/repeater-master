import{b as g}from "./8445cc3c.js";import{f as b,H as C,ad as F,A as N,l as _,j as l,u as e,o as c,i as s,ae as r,af as o,ap as y,e as A,ai as D,n as h,aj as w,ak as x}from "./315a852e.js";import{M as m}from "./index.148f4d09.js";const I={class:"flex justify-end"},U=b({__name:"Add",setup(M){const d=C(null),u=F({name:"",address:"",username:"",password:"",auth_type:"LOGIN"}),i={name:{required:!0,message:"\u8BF7\u8F93\u5165\u540D\u79F0",trigger:"blur"},address:{required:!0,message:"\u8BF7\u8F93\u5165\u670D\u52A1\u5668\u5730\u5740",trigger:"blur"},username:{required:!0,message:"\u8BF7\u8F93\u5165\u7528\u6237",trigger:"blur"},password:{required:!0,message:"\u8BF7\u8F93\u5165\u5BC6\u7801",trigger:"blur"},auth_type:{required:!0,trigger:["blur","change"],message:"\u8BF7\u9009\u62E9\u7C7B\u578B"}},f=[{label:"LOGIN",value:"LOGIN"},{label:"PLAIN",value:"PLAIN"},{label:"CRAM-MD5",value:"CRAM-MD5"}],v=N();function B(){var n;(n=d.value)==null||n.validate(a=>{var t;a?(console.log("errors"),(t=m)==null||t.error("\u9A8C\u8BC1\u5931\u8D25")):g(u).then(k=>{var p;(p=m)==null||p.success("success"),v.push({name:"CredentialStmp"})})})}return(n, a)=>(c(),_(e(x),{title:"\u6DFB\u52A0STMP\u51ED\u8BC1"},{default:l(()=>[s(e(w),{ref_key:"formRef",ref:d,model:u,"label-placement":"left","label-width":160,rules:i,style:{maxWidth:"640px"}},{default:l(()=>[s(e(r),{label:"\u540D\u79F0:",path:"name"},{default:l(()=>[s(e(o),{value:u.name,"onUpdate:value":a[0]||(a[0]= t=>u.name=t)},null,8,["value"])]),_:1}),s(e(r),{label:"\u5730\u5740:",path:"address"},{default:l(()=>[s(e(o),{value:u.address,"onUpdate:value":a[1]||(a[1]= t=>u.address=t)},null,8,["value"])]),_:1}),s(e(r),{label:"\u7528\u6237\u540D:",path:"username"},{default:l(()=>[s(e(o),{value:u.username,"onUpdate:value":a[2]||(a[2]= t=>u.username=t)},null,8,["value"])]),_:1}),s(e(r),{label:"\u5BC6\u7801:",path:"password"},{default:l(()=>[s(e(o),{value:u.password,"onUpdate:value":a[3]||(a[3]= t=>u.password=t)},null,8,["value"])]),_:1}),s(e(r),{label:"\u8BA4\u8BC1\u7C7B\u578B:",path:"auth_type"},{default:l(()=>[s(e(y),{value:u.auth_type,"onUpdate:value":a[4]||(a[4]= t=>u.auth_type=t),placeholder:"\u8BF7\u9009\u62E9",options:f},null,8,["value"])]),_:1}),A("div",I,[s(e(D),{onClick:B,type:"primary"},{default:l(()=>[h("\u63D0\u4EA4")]),_:1})])]),_:1},8,["model"])]),_:1}))}});export{U as default};
