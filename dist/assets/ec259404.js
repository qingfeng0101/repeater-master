import{e as f,i as _}from "./3a8f2c57.js";import{f as g,ad as h,H as k,A as y,c as C,i,j as v,u as n,G as t,av as l,aw as x,ai as o,o as D,n as z,am as F}from "./315a852e.js";import{M as N}from "./index.148f4d09.js";import{_ as b}from "./d304026b.js";const T=g({__name:"index",setup(B){const c=[{title:"\u540D\u79F0",key:"name"},{title:"\u6807\u7B7E\u96C6\u5408\u5217\u8868",key:"label_set_list",render(e){return e.label_set_list.map(a=>t(x,{delay:500,trigger:"hover"},{trigger:()=>t(l,{class:"mb-4px mr-8px",size:"small"},{default:()=>a.service_name||a.namespace||a.instance}),default:()=>JSON.stringify(a)}))}},{title:"\u4E25\u91CD\u7A0B\u5EA6\u5217\u8868",key:"severity_set",render(e){return e.severity_set?e.severity_set.map(a=>t(l,{class:"mb-4px mr-8px",size:"small"},{default:()=>a})):""}},{title:"\u5907\u6CE8",key:"comment"},{title:"\u64CD\u4F5C",key:"actions",width:180,render(e){return[t(o,{size:"small",type:"primary",class:"mr-8px",onClick:()=>u.push({name:"FakePrometheusDatasourceEdit",params:{id:e.id}})},{default:()=>"\u7F16\u8F91"}),t(b,{size:"small",type:"error",onClick:()=>m(e)},{default:()=>"\u5220\u9664"})]}}],s=h({page:1,pageSize:15,showSizePicker:!0,pageSizes:[10,15,25,30,50,100],onChange: e=>{s.page=e},onPageSizeChange: e=>{s.pageSize=e,s.page=1}}),r=k([]);f().then(e=>{r.value=e.data});const u=y();function p(){u.push({name:"FakePrometheusDatasourceAdd"})}function m(e){_(e.id).then(()=>{var a;r.value.splice(r.value.indexOf(e),1),(a=N)==null||a.success("success")})}return(e, a)=>(D(),C("div",null,[i(n(o),{type:"primary",onClick:p,class:"mb-8px"},{default:v(()=>[z("\u6DFB\u52A0")]),_:1}),i(n(F),{data:r.value,pagination:s,columns:c,"row-key": d=>d.id},null,8,["data","pagination","row-key"])]))}});export{T as default};
