const SitemapGenerator = require('sitemap-generator');
const fs = require("fs");
const generator = SitemapGenerator('https://aida.market', {
	maxDepth: 0,
	filepath: '/var/www/aida/shop_for_toma/dist/sitemap.xml',
	maxEntriesPerFile: 50000,
	stripQuerystring: true,
	lastMod:true
});

try {
	fs.unlinkSync('/var/www/aida/shop_for_toma/dist/sitemap.xml');
}catch(e){
	console.log(e)
}
generator.start();
generator.stop()
generator.on('done', () => {
  	try {
		var sitemap = fs.readFileSync(`/var/www/aida/shop_for_toma/dist/sitemap.xml`,'utf8')		
		if(sitemap){
			sitemap = sitemap.replace(/https\:\/\//gi,'https://m.')
			fs.writeFileSync('/var/www/aida/shop_for_toma/dist/m.sitemap.xml',sitemap)
		}
	}catch(e){
		console.log(e)
	}
});