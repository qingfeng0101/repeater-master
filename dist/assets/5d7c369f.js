import{h as c}from "./3a8f2c57.js";import{f as g,ad as h,H as o,l as m,u as f,G as i,av as y,aG as p,o as k,am as v}from "./315a852e.js";const F=g({__name:"index",setup(w){const t=h({page:1,pageSize:15,showSizePicker:!0,pageSizes:[10,15,25,30,50,100],onChange: e=>{t.page=e},onPageSizeChange: e=>{t.pageSize=e,t.page=1}}),d=[{title:"\u73AF\u5883",key:"labels.environment",width:"100",ellipsis:{tooltip:!0}},{title:"\u540D\u79F0",key:"labels.alertname",width:"180",ellipsis:{tooltip:!0}},{title:"\u670D\u52A1\u540D",key:"labels.service_name",width:"240",ellipsis:{tooltip:!0}},{title:"\u7EA7\u522B",key:"labels.severity",width:"80",render(e){var n;const r={high:"#f0a020",critical:"#d03050",warning:"#af52de",info:"#909399"};let a=(n=e.labels)==null?void 0:n.severity;return i(y,{color:{color:r[a]}},{default:()=>{var u;return(u=e.labels)==null?void 0:u.severity}})}},{title:"\u72B6\u6001",key:"status.state",width:"100",ellipsis:{tooltip:!0}},{title:"\u8BE6\u60C5",key:"annotations.description",ellipsis:{tooltip:!0}},{title:"\u5F00\u59CB\u65F6\u95F4",key:"startsAt",width:"170",render(e){return i(p,{time:Date.parse(e.startsAt)})}},{title:"\u7ED3\u675F\u65F6\u95F4",key:"endsAt",width:"170",render(e){return e.endsAt?i(p,{time:Date.parse(e.endsAt)}):"\u2014\u2014"}}],s=o([]),l=o(!0);return c().then(e=>{s.value=e.data}).finally(()=>l.value=!1),(e, r)=>(k(),m(f(v),{data:s.value,pagination:t,columns:d,"row-key": a=>a.id,loading:l.value},null,8,["data","pagination","row-key","loading"]))}});export{F as default};