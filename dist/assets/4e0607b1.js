import{j as d,k as p}from "./8445cc3c.js";import{f as m,ad as f,H as k,A as _,c as g,i as s,j as h,u as i,G as C,o as y,n as x,ai as D,am as v}from "./315a852e.js";import{M as z}from "./index.148f4d09.js";import{_ as S}from "./d304026b.js";const E=m({__name:"index",setup(w){const r=[{title:"\u540D\u79F0",key:"name"},{title:"\u5E94\u7528ID",key:"corp_id"},{title:"\u5E94\u7528Secret",key:"corp_secret"},{title:"\u64CD\u4F5C",key:"actions",width:180,render(e){return[C(S,{size:"small",type:"error",onClick:()=>c(e)},{default:()=>"\u5220\u9664"})]}}],a=f({page:1,pageSize:15,showSizePicker:!0,pageSizes:[10,15,25,30,50,100],onChange: e=>{a.page=e},onPageSizeChange: e=>{a.pageSize=e,a.page=1}}),t=k([]);d().then(e=>{t.value=e.data});const o=_();function u(){o.push({name:"CredentialLarkAdd"})}function c(e){p(e.id).then(()=>{var n;t.value.splice(t.value.indexOf(e),1),(n=z)==null||n.success("success")})}return(e, n)=>(y(),g("div",null,[s(i(D),{type:"primary",onClick:u,class:"mb-8px"},{default:h(()=>[x("\u6DFB\u52A0")]),_:1}),s(i(v),{data:t.value,pagination:a,columns:r,"row-key": l=>l.id},null,8,["data","pagination","row-key"])]))}});export{E as default};