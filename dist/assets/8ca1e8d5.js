import{a as m,b as f}from "./ad89c71c.js";import{_ as g}from "./d304026b.js";import{f as h,ad as k,H as _,A as y,c as S,i as u,j as C,u as r,G as t,av as x,ai as l,o as z,n as D,am as v}from "./315a852e.js";const A=h({__name:"index",setup(F){const c=[{title:"\u540D\u79F0",key:"name"},{title:"\u6D88\u606F\u53D1\u9001\u5668",key:"sets",render(e){return!(e!=null&&e.sets)||e.sets.map(i=>t(x,{size:"small",type:"warning",class:"mr-8px"},{default:()=>i}))}},{title:"\u64CD\u4F5C",key:"actions",width:180,render(e){return[t(l,{size:"small",type:"primary",class:"mr-8px",onClick:()=>n.push({name:"SenderSetsEdit",params:{id:e.id}})},{default:()=>"\u7F16\u8F91"}),t(g,{size:"small",type:"error",onClick:()=>o(e)},{default:()=>"\u5220\u9664"})]}}],a=k({page:1,pageSize:15,showSizePicker:!0,pageSizes:[10,15,25,30,50,100],onChange: e=>{a.page=e},onPageSizeChange: e=>{a.pageSize=e,a.page=1}}),s=_([]);m().then(e=>{s.value=e.data});const n=y();function d(){n.push({name:"SenderSetsAdd"})}function o(e){f(e.id).then(()=>{s.value.splice(s.value.indexOf(e),1),Message==null||Message.success("success")})}return(e, i)=>(z(),S("div",null,[u(r(l),{class:"mb-8px",type:"primary",onClick:d},{default:C(()=>[D("\u6DFB\u52A0")]),_:1}),u(r(v),{data:s.value,pagination:a,columns:c,"row-key": p=>p.id},null,8,["data","pagination","row-key"])]))}});export{A as default};
