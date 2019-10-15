const SitemapGenerator = require('sitemap-generator');
const cron = require('node-schedule');

const generator = SitemapGenerator('https://aida.market', {
  	maxDepth: 0,
  	filepath: '/var/www/aida/shop_for_toma/dist/sitemap.xml',
  	maxEntriesPerFile: 50000,
  	stripQuerystring: true,
  	lastMod:true
});
generator.start();
generator.stop();