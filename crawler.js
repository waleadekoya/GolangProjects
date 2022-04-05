const axios = require("axios");
const cheerio = require('cheerio');
const fs = require('fs');
const csvjson = require('csvjson');


class Utils {
    static loadHtml = async url => {
        const response = await axios.get(url);
        const html = response.data;
        return cheerio.load(html); // https://cheerio.js.org/
    }
    static logErrors = err => {
        if (err) {
            console.log(err);
            throw new Error(err);
        }
    }
    static writeJson = results => {
        console.log(`Writing results to JSON file...`)
        const saveJson = JSON.stringify(results, null, 4);
        fs.writeFile('jobs.json', saveJson, 'utf-8', (err) => {
            Utils.logErrors(err);
        })
        console.log(`Write to JSON file completed.`)
    };
    static writeToCSV = () => {
        console.log(`Converting JSON file to CSV...`);
        fs.readFile('jobs.json', 'utf-8',
            (err, fileContent) => {
                Utils.logErrors(err);
                const csvData = csvjson.toCSV(fileContent, {headers: 'key'});
                fs.writeFile('jobs.csv', csvData, 'utf-8',
                    (err) => Utils.logErrors(err))
                console.log('Success!');
            })
    };
}

class ReedJobsCrawler {

    constructor(searchKey, startingSalary = 100000, isContract = 1) {
        this._searchKey = searchKey.split(" ").join("-");
        this._startingSalary = startingSalary;
        this._isContract = isContract ? "True" : "False";
        this.getJobs().then(r => r)
    }

    handlePagination = async () => {
        let baseUrl = `https://www.reed.co.uk/jobs/${this._searchKey}-jobs?pageno=1&salaryfrom=${this._startingSalary}&contract=${this._isContract}&datecreatedoffset=LastThreeDays`
        const $ = Utils.loadHtml(baseUrl);
        let resultsCount = (await $)("#content > div.row.search-results > div.col-sm-11.col-xs-12.page-title > span").text()
        console.log(`Found a total of ${resultsCount ? parseInt(resultsCount) : 0} results.`)
        return resultsCount ? Math.ceil(parseInt(resultsCount) / 25) : 1;
    }
    getJobs = async () => {
        let jobDetails = []
        try {
            let pageCount = await this.handlePagination()
            if (pageCount) {
                for (let i = 1; i < pageCount + 1; i++) {
                    let jobsCount = jobDetails.length;
                    console.log(`Crawling page ${i}...`);
                    let url = `https://www.reed.co.uk/jobs/${this._searchKey}-jobs?pageno=${i}&salaryfrom=${this._startingSalary}&contract=${this._isContract}&datecreatedoffset=LastThreeDays`
                    console.log(url);
                    const $ = Utils.loadHtml(url);
                    const jobsInfo = (await $)(".col-sm-12.col-md-9.col-lg-9.details > header > h3 > a");
                    await this.extractJobInfo(jobsInfo, jobDetails);
                    jobsCount = jobDetails.length - jobsCount;
                    console.log(`Successfully retrieved ${jobsCount} jobs on page ${i}`)
                }
                Utils.writeJson(jobDetails)
                Utils.writeToCSV();
            }
        } catch (error) {
            console.log(error)
        }
    };
    extractJobInfo = async (jobsInfo, jobDetails) => {
        for (let value of jobsInfo) {
            let jobUrl = "https://www.reed.co.uk/" + value.attribs.href
            const $ = Utils.loadHtml(jobUrl)
            let title = (await $)("#content > div.job-details.row > div.col-xs-12.col-sm-12.col-md-9 > article > div > header > div.col-xs-12 > h1").text(),
                advertiser = (await $)("#content > div.job-details.row > div.col-xs-12.col-sm-12.col-md-9 > article > div > header > div.col-xs-12 > div > span > a > span").text(),
                rate = (await $)("#content > div.job-details.row > div.col-xs-12.col-sm-12.col-md-9 > article > div > div.description-container.row > div > div.hidden-xs > div > div.salary.col-xs-12.col-sm-6.col-md-6.col-lg-6 > span > span:nth-child(2)").text(),
                location = (await $)("#content > div.job-details.row > div.col-xs-12.col-sm-12.col-md-9 > article > div > div.description-container.row > div > div.hidden-xs > div > div.location.col-xs-12.col-sm-6.col-md-6.col-lg-6 > span:nth-child(3) > span").text(),
                jobText = (await $)("span[itemprop='description']").text(),
                isRemote = jobText.toLowerCase().includes("remote"),
                job = {title, isRemote, rate, location, advertiser, jobUrl};
            jobDetails.push(job)
        }
    };
}

new ReedJobsCrawler("python developer", 100000, 1)

const getPageCount = async url => {
    const $ = Utils.loadHtml(url);
    let resultsCount = (await $)("span.at-facet-header-total-results")['0'].children[0].data;
    return resultsCount ? Math.ceil(parseInt(resultsCount) / 25) : 0
}


const GetUrls = async () => {
    let urls = [], baseUrl = "https://www.totaljobs.com/jobs/contract/remote-python?postedWithin=3";
    const pageCount = await getPageCount(baseUrl);
    console.log(pageCount)
    for (let i = 1; i < pageCount + 1; i++) {
        let dynamicUrl = `https://www.totaljobs.com/jobs/contract/remote-python?page=${i}&postedWithin=3`
        console.log(dynamicUrl)
        const $ = Utils.loadHtml(dynamicUrl);
        let pageInfo = (await $)("div.ResultsContainer-sc-1rtv0xy-2.dmSilN > div")
        console.log(pageInfo)
        for (const value of pageInfo) {
            urls.push(value.attribs.href)
            console.log(value.attribs.href)
        }
    }
    console.log(urls)
    console.log(urls.length)
}
// https://www.totaljobs.com/jobs/contract/python-developer?salary=100000&salarytypeid=1&postedWithin=3

// GetUrls() //.then(r => r)

// #app-unifiedResultlist-6a42d577-ca54-4f8e-9cfd-53223dd1e449 > div > div.container > div > div.ResultsWrapper-sc-h54hyq-2.cjTGHz.col-lg-9 > div.ResultsContainer-sc-1rtv0xy-2.dmSilN > div

