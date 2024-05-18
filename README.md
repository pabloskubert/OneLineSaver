### OneLineSaver 

Save your one-liners for later reuse. 

#### Usage 
-> olins "commands... {var1} {var2} {var3}" name

Examples:  

`olins "subfinder -d {site} -silent | dnsx -silent | cut -d ' ' -f1  | grep --color 'api\|dev\|stg\|admin'" subfinder_site `  

Then use:
`subfinder_site example.com`  

Scan IPs:
`olins "cat {ips_file} | xargs -L 100 shodan scan submit --wait 0" shodan_scan_ips`

Then use:
`shodan_scan_ips ips.txt`

Add ~/.olins to $PATH
