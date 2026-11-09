const api = async (path, options={}) => {
  const response = await fetch(path, {headers:{"Content-Type":"application/json"}, ...options});
  if (!response.ok) throw new Error(await response.text() || "Fehler");
  return response.status === 204 ? null : response.json();
};
const el = id => document.getElementById(id);
const formData = form => Object.fromEntries(new FormData(form).entries());
const esc = value => String(value).replace(/[&<>"']/g, c => ({"&":"&amp;","<":"&lt;",">":"&gt;",'"':"&quot;","'":"&#39;"}[c]));
let skis=[], persons=[];
function table(headers, rows) {
  if (!rows.length) return "<p>Noch keine Einträge vorhanden.</p>";
  return `<table><thead><tr>${headers.map(h=>`<th>${h}</th>`).join("")}</tr></thead><tbody>${rows.join("")}</tbody></table>`;
}
async function refresh() {
  const [dash,s,p,l,r] = await Promise.all([api("/api/dashboard"),api("/api/skis"),api("/api/persons"),api("/api/loans"),api("/api/reservations")]);
  skis=s; persons=p;
  el("available").textContent=dash.availableSkis; el("rented").textContent=dash.rentedSkis; el("total").textContent=dash.totalSkis; el("people").textContent=dash.persons;
  el("skiList").innerHTML=table(["ID","Name","Kategorie","Größe","Preis/Tag","Status",""],s.map(x=>`<tr><td>${x.id}</td><td>${esc(x.name)}</td><td>${esc(x.category)}</td><td>${x.size} cm</td><td>€ ${Number(x.dailyPrice).toFixed(2)}</td><td>${esc(x.status)}</td><td><button class="delete" onclick="deleteSki(${x.id})">Löschen</button></td></tr>`));
  el("personList").innerHTML=table(["ID","Name","E-Mail",""],p.map(x=>`<tr><td>${x.id}</td><td>${esc(x.name)}</td><td>${esc(x.email)}</td><td><button class="delete" onclick="deletePerson(${x.id})">Löschen</button></td></tr>`));
  const available=s.filter(x=>x.status==="available");
  const opts=arr=>arr.map(x=>`<option value="${x.id}">${esc(x.name)} (#${x.id})</option>`).join("");
  el("skiSelect").innerHTML=opts(available); el("reserveSki").innerHTML=opts(s);
  el("personSelect").innerHTML=opts(p); el("reservePerson").innerHTML=opts(p);
  el("loanList").innerHTML=table(["ID","Ski-ID","Person-ID","Datum","Tage","Preis/Tag","Gesamt"],l.map(x=>`<tr><td>${x.id}</td><td>${x.skiId}</td><td>${x.personId}</td><td>${esc(x.date)}</td><td>${x.days}</td><td>€ ${Number(x.dailyPrice).toFixed(2)}</td><td>€ ${Number(x.totalPrice).toFixed(2)}</td></tr>`));
  el("reservationList").innerHTML=table(["ID","Ski-ID","Person-ID","Datum","Status"],r.map(x=>`<tr><td>${x.id}</td><td>${x.skiId}</td><td>${x.personId}</td><td>${esc(x.date)}</td><td>${esc(x.status)}</td></tr>`));
}
async function submitForm(id,path,convert) {
  const form=el(id); form.addEventListener("submit",async e=>{e.preventDefault();try{const data=formData(form);if(convert)convert(data);await api(path,{method:"POST",body:JSON.stringify(data)});form.reset();await refresh();}catch(err){alert(err.message);}});
}
submitForm("skiForm","/api/skis",d=>{d.size=Number(d.size);d.dailyPrice=Number(d.dailyPrice);});
submitForm("personForm","/api/persons");
submitForm("loanForm","/api/loans",d=>{d.skiId=Number(d.skiId);d.personId=Number(d.personId);d.days=Number(d.days);});
submitForm("reservationForm","/api/reservations",d=>{d.skiId=Number(d.skiId);d.personId=Number(d.personId);});
window.deleteSki=async id=>{if(confirm("Ski wirklich löschen?")){await api(`/api/skis/${id}`,{method:"DELETE"});refresh();}};
window.deletePerson=async id=>{if(confirm("Person wirklich löschen?")){await api(`/api/persons/${id}`,{method:"DELETE"});refresh();}};
refresh().catch(err=>alert("Backend nicht erreichbar: "+err.message));
