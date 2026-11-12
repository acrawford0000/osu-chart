import test from 'node:test';
import assert from 'node:assert/strict';
import { get } from 'svelte/store';
import { createPlayerState } from '../src/lib/playerState.js';
import { createAuthState } from '../src/lib/authState.js';
import { normalizeError } from '../src/lib/errors.js';

const deferred = () => { let resolve, reject; const promise = new Promise((a,b) => {resolve=a;reject=b}); return {promise,resolve,reject}; };
const profile = name => ({id:1,username:name});
const flush = () => new Promise(resolve => setTimeout(resolve,0));

test('loading is derived and stays balanced after empty refresh, failure, and retry', async () => {
 let fail=true;const messages=[];
 const state=createPlayerState({GetUser:async name=>profile(name),GetStatsUpdates:async()=>{if(fail)throw 'NETWORK_TIMEOUT: History timed out. Please retry.';return []}},message=>messages.push(message));
 await state.fetchPlayerStats();assert.equal(get(state.loadingCount),0);
 assert.equal(await state.addPlayer(' demo '),false);
 assert.equal(get(state.loadingCount),0);assert.equal(get(state.isLoading),false);
 assert.equal(get(state.players)[0].status,'error');assert.match(messages[0],/demo: History timed out/);
 fail=false;await state.retryPlayer('demo');
 assert.equal(get(state.players)[0].status,'empty');assert.deepEqual(get(state.players)[0].stats,[]);
 assert.equal(get(state.loadingCount),0);
});
test('duplicate pending adds are rejected and mode is passed to both requests', async () => {
 const pending=deferred(),calls=[];
 const state=createPlayerState({GetUser:async(name,mode)=>{calls.push(['profile',mode]);await pending.promise;return profile(name)},GetStatsUpdates:async(id,mode)=>{calls.push(['history',mode]);return []}});
 const first=state.addPlayer('Demo','mania');
 assert.equal(await state.addPlayer(' demo ','standard'),false);assert.equal(get(state.loadingCount),1);
 pending.resolve();await first;
 assert.deepEqual(calls,[['profile','mania'],['history','mania']]);assert.equal(get(state.players).length,1);
});
test('late mode responses cannot replace newer data; cache is keyed by player and mode', async () => {
 const mania=deferred();const calls=[];
 const state=createPlayerState({GetUser:async name=>profile(name),GetStatsUpdates:async(id,mode)=>{calls.push([id,mode]);return mode==='mania'?mania.promise:[{mode}]}});
 await state.addPlayer('demo');
 const older=state.changePlayerMode('demo','mania');
 assert.equal(get(state.players)[0].stats,null);
 await state.changePlayerMode('demo','taiko');
 mania.resolve([{mode:'mania'}]);await older;
 assert.equal(get(state.players)[0].selectedMode,'taiko');assert.deepEqual(get(state.players)[0].stats,[{mode:'taiko'}]);
 await state.changePlayerMode('demo','standard');
 assert.deepEqual(get(state.players)[0].stats,[{mode:'standard'}]);
 assert.equal(calls.filter(c=>c[1]==='standard').length,1);
});
test('remove and clear invalidate pending requests without resurrecting players', async () => {
 for (const action of ['remove','clear']) {
  const pending=deferred();const state=createPlayerState({GetUser:async name=>{await pending.promise;return profile(name)},GetStatsUpdates:async()=>[]});
  const load=state.addPlayer('demo');
  if(action==='remove')state.removePlayer('demo');else state.clearPlayers();
  pending.resolve();await load;
  assert.deepEqual(get(state.players),[]);assert.equal(get(state.isLoading),false);
 }
});
test('a removed and re-added player owns a new request even with the same name', async()=>{
 const first=deferred(),second=deferred();let call=0;
 const state=createPlayerState({GetUser:async name=>{await (++call===1?first.promise:second.promise);return profile(name)},GetStatsUpdates:async()=>[]});
 const oldLoad=state.addPlayer('demo');state.removePlayer('demo');const newLoad=state.addPlayer('demo');
 first.resolve();await oldLoad;assert.equal(get(state.loadingCount),1);
 second.resolve();await newLoad;assert.equal(get(state.players)[0].status,'empty');assert.equal(get(state.loadingCount),0);
});
test('profile errors remain retryable and never leave a spinner running',async()=>{
 let fail=true;
 const state=createPlayerState({GetUser:async name=>{if(fail)throw new Error('NOT_FOUND: No player found.');return profile(name)},GetStatsUpdates:async()=>[{value:1}]});
 await state.addPlayer('demo');assert.equal(get(state.players)[0].error.code,'NOT_FOUND');assert.equal(get(state.isLoading),false);
 fail=false;await state.retryPlayer('demo');assert.equal(get(state.players)[0].status,'ready');
});
test('parallel players have independent modes and loading states',async()=>{
 const one=deferred(),two=deferred();
 const state=createPlayerState({GetUser:async name=>({id:name==='one'?1:2,username:name}),GetStatsUpdates:async(id,mode)=>{assert.equal(mode,id===1?'mania':'taiko');return id===1?one.promise:two.promise}});
 const a=state.addPlayer('one','mania'),b=state.addPlayer('two','taiko');await flush();assert.equal(get(state.loadingCount),2);
 one.resolve([]);await a;assert.equal(get(state.loadingCount),1);
 two.resolve([{value:2}]);await b;assert.equal(get(state.loadingCount),0);
});
const authAPI=overrides=>({GetCredentials:async()=>({client_id:'1',client_secret:'fake-only'}),IsClientValid:async()=>true,CreateNewClient:async()=>{},SaveOsuAuthCredentials:async()=>{},...overrides});
test('authentication restores existing clients, distinguishes missing credentials and failures',async()=>{
 const ready=createAuthState(authAPI());await ready.initialize();assert.equal(get(ready.clientValid),true);assert.equal(get(ready.credentialsSet),true);
 const missing=createAuthState(authAPI({GetCredentials:async()=>({client_id:'',client_secret:''})}));await missing.initialize();assert.equal(get(missing.authState).status,'missing');
 const broken=createAuthState(authAPI({GetCredentials:async()=>{throw 'CONFIG_ERROR: Could not read settings.'}}));await broken.initialize();assert.equal(get(broken.authState).error.code,'CONFIG_ERROR');assert.equal(get(broken.clientValid),false);
});
test('credentials can be edited after a valid login; failed saves keep the previous ready state',async()=>{
 let calls=0,fail=false;
 const auth=createAuthState(authAPI({SaveOsuAuthCredentials:async(id,secret)=>{calls++;assert.equal(id,'2');assert.equal(secret,'updated-fake');if(fail)throw 'AUTH_INVALID: Rejected credentials.'}}));
 await auth.initialize();auth.clientId.set('2');auth.clientSecret.set('updated-fake');assert.equal(await auth.saveCredentials(),true);assert.equal(calls,1);
 fail=true;await assert.rejects(auth.saveCredentials());assert.equal(get(auth.clientValid),true);assert.equal(get(auth.authState).error.code,'AUTH_INVALID');
});
test('saving credentials prevents duplicate submission and validates before calling the API',async()=>{
 let calls=0;const pending=deferred();
 const auth=createAuthState(authAPI({SaveOsuAuthCredentials:async()=>{calls++;await pending.promise}}));await auth.initialize();
 auth.clientId.set('invalid');await assert.rejects(auth.saveCredentials());assert.equal(calls,0);
 auth.clientId.set('2');const save=auth.saveCredentials();assert.equal(await auth.saveCredentials(),false);assert.equal(calls,1);await auth.initialize();assert.equal(get(auth.authState).status,'saving');pending.resolve();await save;assert.equal(get(auth.clientValid),true);
});
test('normalizes Wails string errors, JavaScript errors, and unknown failures',()=>{
 assert.deepEqual(normalizeError('NETWORK_TIMEOUT: Try again.'),{code:'NETWORK_TIMEOUT',message:'Try again.'});
 assert.equal(normalizeError(new Error('Readable message')).message,'Readable message');assert.equal(normalizeError(null).code,'UNKNOWN');
});

test('history cache expires so switching back to an older mode refreshes data',async t=>{
 let time=1000,calls=0;t.mock.method(Date,'now',()=>time);
 const state=createPlayerState({GetUser:async name=>profile(name),GetStatsUpdates:async()=>[{value:++calls}]});
 await state.addPlayer('demo');await state.changePlayerMode('demo','taiko');await state.changePlayerMode('demo','standard');assert.equal(calls,2);
 time+=5*60*1000+1;await state.changePlayerMode('demo','taiko');await state.changePlayerMode('demo','standard');assert.equal(calls,4);
});
test('saving is blocked during initialization without invalidating the pending startup',async()=>{
 const pending=deferred();let calls=0;
 const auth=createAuthState(authAPI({GetCredentials:async()=>{await pending.promise;return {client_id:'1',client_secret:'fake-only'}},SaveOsuAuthCredentials:async()=>{calls++}}));
 const initialize=auth.initialize();auth.clientId.set('2');auth.clientSecret.set('draft');assert.equal(await auth.saveCredentials(),false);
 pending.resolve();await initialize;assert.equal(calls,0);assert.equal(get(auth.authState).status,'ready');
});
