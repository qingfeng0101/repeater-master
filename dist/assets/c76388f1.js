import{f as x,H as o,W as g,l as y,j as h,u as v,o as c,e as C,D as i,an as b,x as k,c as D,y as E,ax as N,G as t,ay as n,av as f,az as S}from "./315a852e.js";import{a as z,d as B}from "./b4128dcd.js";const w=["onClick"],F=x({__name:"index",setup(T){const u=o({}),d={nodeWidth:120,nodeHeight:60,levelHeight:130},s=o();z().then(a=>{s.value=a.data}),B().then(a=>{u.value=a.data});const p=g();function m(a){console.log(s.value[a]);let l=s.value[a];!l||p.create({title:a,positiveText:"\u786E\u5B9A",content:()=>t(S,{column:1,labelPlacement:"left"},{default:()=>[t(n,{label:"\u901A\u77E5\u76EE\u6807"},{default:()=>l.Contacts.map(e=>t(f,{size:"small",class:"mr-4px"},{default:()=>e}))}),t(n,{label:"\u53D1\u9001\u5668"},{default:()=>l.SenderSet.map(e=>t(f,{size:"small",class:"mr-4px"},{default:()=>e}))}),t(n,{label:"\u53D1\u9001\u5668\u8FC7\u6EE4"},{default:()=>JSON.stringify(l.SenderFilter)})]})})}return(a, l)=>(c(),y(v(N),{class:"w-full h-full bg-#fff break-all",dataset:u.value,config:d},{node:h(({node:e,collapsed:_})=>{var r;return[C("span",{class:b(["bg-blue p-2px rd-t-5% w-full",{"text-center":((r=e.key)==null?void 0:r.indexOf("="))===-1}]),style:k({border:_?"1px solid grey":""})},i(e.key),7),e.value?(c(),D("span",{key:0,onClick: H=>m(e.value),class:"bg-#e2e2e2 w-full text-center"},i(e.value),9,w)):E("",!0)]}),_:1},8,["dataset"]))}});export{F as default};