"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[182],{8214:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>a,contentTitle:()=>r,default:()=>l,frontMatter:()=>s,metadata:()=>c,toc:()=>h});var i=t(5893),o=t(1151);const s={},r="OAuth Public Frontend Configuration",c={id:"guides/self-hosting/oauth/configuring-oauth",title:"OAuth Public Frontend Configuration",description:"As of v0.4.7, zrok includes OAuth integration for both Google and GitHub for zrok access public public frontends.",source:"@site/../docs/guides/self-hosting/oauth/configuring-oauth.md",sourceDirName:"guides/self-hosting/oauth",slug:"/guides/self-hosting/oauth/configuring-oauth",permalink:"/docs/guides/self-hosting/oauth/configuring-oauth",draft:!1,unlisted:!1,editUrl:"https://github.com/openziti/zrok/blob/main/docs/../docs/guides/self-hosting/oauth/configuring-oauth.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"OAuth",permalink:"/docs/category/oauth"},next:{title:"Instance Config",permalink:"/docs/guides/self-hosting/instance-configuration"}},a={},h=[{value:"Planning for the OAuth Frontend",id:"planning-for-the-oauth-frontend",level:2},{value:"Configuring a Google OAuth Client ID",id:"configuring-a-google-oauth-client-id",level:2},{value:"OAuth Content Screen",id:"oauth-content-screen",level:3},{value:"Create the OAuth 2.0 Client ID",id:"create-the-oauth-20-client-id",level:3},{value:"Configuring a GitHub Client ID",id:"configuring-a-github-client-id",level:2},{value:"Configuring your Public Frontend",id:"configuring-your-public-frontend",level:2},{value:"Enabling OAuth on a Public Share",id:"enabling-oauth-on-a-public-share",level:2}];function d(e){const n={code:"code",h1:"h1",h2:"h2",h3:"h3",img:"img",p:"p",pre:"pre",...(0,o.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.h1,{id:"oauth-public-frontend-configuration",children:"OAuth Public Frontend Configuration"}),"\n",(0,i.jsxs)(n.p,{children:["As of ",(0,i.jsx)(n.code,{children:"v0.4.7"}),", ",(0,i.jsx)(n.code,{children:"zrok"})," includes OAuth integration for both Google and GitHub for ",(0,i.jsx)(n.code,{children:"zrok access public"})," public frontends."]}),"\n",(0,i.jsx)(n.p,{children:"This integration allows you to create public shares and request that the public frontend authenticate your users against either the Google or GitHub OAuth endpoints (using the user's Google or GitHub accounts). Additionally, you can restrict the email address domain associated with the count to a list of domains that you provide when you create the share."}),"\n",(0,i.jsxs)(n.p,{children:["This is a first step towards a more comprehensive portfolio of user authentication strategies in future ",(0,i.jsx)(n.code,{children:"zrok"})," releases."]}),"\n",(0,i.jsx)(n.h2,{id:"planning-for-the-oauth-frontend",children:"Planning for the OAuth Frontend"}),"\n",(0,i.jsx)(n.p,{children:"The current implementation of the OAuth public frontend uses a HTTP listener to handle redirects from OAuth providers. You'll need to configure a DNS name and a port for this listener that is accessible by your end users. We'll refer to this listener as the \"OAuth frontend\" in this guide."}),"\n",(0,i.jsx)(n.p,{children:'We\'ll use the public DNS address of the OAuth frontend when creating the Google and GitHub OAuth clients below. This address is typically configured into these clients as the "redirect URL" where these clients will send the authenticated users after authentication.'}),"\n",(0,i.jsxs)(n.p,{children:["The ",(0,i.jsx)(n.code,{children:"zrok"})," OAuth frontend will capture the successful authentication and forward the user back to their original destination."]}),"\n",(0,i.jsx)(n.h2,{id:"configuring-a-google-oauth-client-id",children:"Configuring a Google OAuth Client ID"}),"\n",(0,i.jsx)(n.h3,{id:"oauth-content-screen",children:"OAuth Content Screen"}),"\n",(0,i.jsx)(n.p,{children:'Before you can configure an OAuth Client ID in Google Cloud, you have to configure the "OAuth content screen".'}),"\n",(0,i.jsxs)(n.p,{children:["In the Google Cloud console, navigate to: ",(0,i.jsx)(n.code,{children:"APIs & Services > Credentials > OAuth content screen"})]}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(7199).Z+"",width:"1469",height:"1141"})}),"\n",(0,i.jsxs)(n.p,{children:["Here you can give your ",(0,i.jsx)(n.code,{children:"zrok"})," public frontend an identity and branding to match your deployment."]}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(2795).Z+"",width:"1469",height:"1141"})}),"\n",(0,i.jsx)(n.p,{children:"Describe what domains are authorized to access your public frontend and establish contact information."}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(3923).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsxs)(n.p,{children:["Add a non-sensitive scope for ",(0,i.jsx)(n.code,{children:"../auth/userinfo.email"}),". This is important as it allows the ",(0,i.jsx)(n.code,{children:"zrok"})," OAuth frontend to receive the email address of the authenticated user."]}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(2575).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(72).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsx)(n.p,{children:"Now your OAuth content screen is configured."}),"\n",(0,i.jsx)(n.h3,{id:"create-the-oauth-20-client-id",children:"Create the OAuth 2.0 Client ID"}),"\n",(0,i.jsx)(n.p,{children:"Next we create the OAuth Client ID for your public frontend."}),"\n",(0,i.jsxs)(n.p,{children:["In the Google Cloud Console, navigate to: ",(0,i.jsx)(n.code,{children:"APIs & Services > Credentials > + Create Credentials"})]}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(5180).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsxs)(n.p,{children:["Select ",(0,i.jsx)(n.code,{children:"OAuth client ID"})," from the ",(0,i.jsx)(n.code,{children:"+ Create Credentials"})," dropdown."]}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(5871).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsxs)(n.p,{children:["Application type is ",(0,i.jsx)(n.code,{children:"Web Application"}),"."]}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(9572).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsxs)(n.p,{children:['The most important bit here is the "Authorized redirect URIs". You\'re going to want to put a URL here that matches the ',(0,i.jsx)(n.code,{children:"zrok"})," OAuth frontend address that you configured at the start of this guide, but at the end of the URL you're going to append ",(0,i.jsx)(n.code,{children:"/google/oauth"})," to the URL."]}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(1115).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsxs)(n.p,{children:["Save the client ID and the client secret. You'll configure these into your ",(0,i.jsx)(n.code,{children:"frontend.yml"}),"."]}),"\n",(0,i.jsx)(n.p,{children:"With this your Google OAuth client should be configured and ready."}),"\n",(0,i.jsx)(n.h2,{id:"configuring-a-github-client-id",children:"Configuring a GitHub Client ID"}),"\n",(0,i.jsx)(n.p,{children:"Register a new OAuth application through the GitHub settings for the account that owns the application."}),"\n",(0,i.jsxs)(n.p,{children:["Navigate to:",(0,i.jsx)(n.code,{children:"Settings > Developer Settings > OAuth Apps > Register a new application"})]}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(9980).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(4861).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsxs)(n.p,{children:['The "Authorized callback URL" should be configured to match the OAuth frontend address you configured at the start of this guide, with ',(0,i.jsx)(n.code,{children:"/github/oauth"})," appended to the end."]}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(5758).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsx)(n.p,{children:"Create a new client secret."}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.img,{src:t(1995).Z+"",width:"1469",height:"1179"})}),"\n",(0,i.jsxs)(n.p,{children:["Save the client ID and the client secret. You'll configure these into your ",(0,i.jsx)(n.code,{children:"frontend.yml"}),"."]}),"\n",(0,i.jsx)(n.h2,{id:"configuring-your-public-frontend",children:"Configuring your Public Frontend"}),"\n",(0,i.jsxs)(n.p,{children:["The public frontend configuration includes a new ",(0,i.jsx)(n.code,{children:"oauth"})," section:"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-yaml",children:'oauth:\n  bind_address:                   0.0.0.0:8181\n  redirect_url:                   https://oauth.zrok.io\n  cookie_domain:                  zrok.io\n  hash_key:                       "the quick brown fox jumped over the lazy dog"\n  providers:\n    - name:                       google\n      client_id:                  "<client id from google>"\n      client_secret:              "<client secret from google>"\n    - name:                       github\n      client_id:                  "<client id from github>"\n      client_secret:              "<client secret from github>"\n      \n'})}),"\n",(0,i.jsxs)(n.p,{children:["The ",(0,i.jsx)(n.code,{children:"bind_address"})," parameter determines where the OAuth frontend will bind. Should be in ",(0,i.jsx)(n.code,{children:"ip:port"})," format."]}),"\n",(0,i.jsxs)(n.p,{children:["The ",(0,i.jsx)(n.code,{children:"redirect_url"})," parameter determines the base URL where OAuth frontend requests will be redirected."]}),"\n",(0,i.jsxs)(n.p,{children:[(0,i.jsx)(n.code,{children:"cookie_domain"})," is the domain where authentication cookies should be stored."]}),"\n",(0,i.jsxs)(n.p,{children:[(0,i.jsx)(n.code,{children:"hash_key"})," is a unique string for your installation that is used to secure the authentication payloads for your public frontend."]}),"\n",(0,i.jsxs)(n.p,{children:[(0,i.jsx)(n.code,{children:"providers"})," is a list of configured providers for this public frontend. The current implementation supports ",(0,i.jsx)(n.code,{children:"google"})," and ",(0,i.jsx)(n.code,{children:"github"})," as options."]}),"\n",(0,i.jsxs)(n.p,{children:["Both the ",(0,i.jsx)(n.code,{children:"google"})," and ",(0,i.jsx)(n.code,{children:"github"})," providers accept a ",(0,i.jsx)(n.code,{children:"client_id"})," and ",(0,i.jsx)(n.code,{children:"client_secret"})," parameter. These values are provided when you configure the OAuth clients at Google or GitHub."]}),"\n",(0,i.jsx)(n.h2,{id:"enabling-oauth-on-a-public-share",children:"Enabling OAuth on a Public Share"}),"\n",(0,i.jsx)(n.p,{children:"With your public frontend configured to support OAuth, you can test this by creating a public share. There are new command line options to support this:"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-text",children:'$ zrok share public\nError: accepts 1 arg(s), received 0\nUsage:\n  zrok share public <target> [flags]\n\nFlags:\n  -b, --backend-mode string               The backend mode {proxy, web, caddy, drive} (default "proxy")\n      --basic-auth stringArray            Basic authentication users (<username:password>,...)\n      --frontends stringArray             Selected frontends to use for the share (default [public])\n      --headless                          Disable TUI and run headless\n  -h, --help                              help for public\n      --insecure                          Enable insecure TLS certificate validation for <target>\n      --oauth-check-interval duration     Maximum lifetime for OAuth authentication; reauthenticate after expiry (default 3h0m0s)\n      --oauth-email-domains stringArray   Allow only these email domains to authenticate via OAuth\n      --oauth-provider string             Enable OAuth provider [google, github]\n\nGlobal Flags:\n  -p, --panic     Panic instead of showing pretty errors\n  -v, --verbose   Enable verbose logging\n'})}),"\n",(0,i.jsxs)(n.p,{children:["The ",(0,i.jsx)(n.code,{children:"--oauth-provider"})," flag enables OAuth for the share using the specified provider."]}),"\n",(0,i.jsxs)(n.p,{children:["The ",(0,i.jsx)(n.code,{children:"--oauth-email-domains"})," flag accepts a comma-separated list of authenticated email address domains that are allowed to access the share."]}),"\n",(0,i.jsxs)(n.p,{children:["The ",(0,i.jsx)(n.code,{children:"--oauth-check-interval"})," flag specifies how frequently the authentication must be checked."]}),"\n",(0,i.jsx)(n.p,{children:"An example public share:"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-text",children:"zrok share public --backend-mode web --oauth-provider github --oauth-email-domains zrok.io ~/public\n"})})]})}function l(e={}){const{wrapper:n}={...(0,o.a)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},9980:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/github_create_oauth_application_1-dbb289a694d0c99b50cb949654d818f8.png"},4861:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/github_create_oauth_application_2-bc9d4c3f25853d608870eb220d00e5ee.png"},5758:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/github_create_oauth_application_3-7973d63cd117eaba72fbaeb4ff119a39.png"},1995:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/github_create_oauth_application_4-6fed398013c9e6c3a31e5721adac4a4c.png"},5180:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/google_create_credentials_1-e61ee7e8fa51bdc93feab84235a90673.png"},5871:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/google_create_credentials_2-60cf4edb52f453d605907c17400e0800.png"},9572:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/google_create_credentials_3-1b3db3f9057d8626d64c5466dbb05ec7.png"},1115:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/google_create_credentials_4-f720031df1d09f997a18842745e9ea0d.png"},7199:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/google_oauth_content_screen_2-b94a6456ce9b13e053b4c07d8f233e84.png"},2795:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/google_oauth_content_screen_3-edf62ea2b291b58093289d31c2dae58d.png"},3923:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/google_oauth_content_screen_4-4bc7e07b06c8a9a3c1e8f766f6f1c5a6.png"},2575:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/google_oauth_content_screen_5-7c375cf49d8e2e392ca12b584462ab59.png"},72:(e,n,t)=>{t.d(n,{Z:()=>i});const i=t.p+"assets/images/google_oauth_content_screen_6-9b0fe216a782ef378313650e99ea52a1.png"},1151:(e,n,t)=>{t.d(n,{Z:()=>c,a:()=>r});var i=t(7294);const o={},s=i.createContext(o);function r(e){const n=i.useContext(s);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function c(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:r(e.components),i.createElement(s.Provider,{value:n},e.children)}}}]);