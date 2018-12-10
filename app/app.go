package main

import (
"fmt"
"log"
"net/http"
"os"
)

var cdn string = "http://lawas-static.s3-website-us-east-1.amazonaws.com"
var document string = `<!DOCTYPE html>
<html lang="en">
	<head>
		<title>Lawas.me - Project Site</title>
		<meta name="keywords" content="gilberto lawas php web developer apache2 mysql zend cakephp symfony2 doctrine">
		<meta name="description" content="Gilberto Lawas Project Site , Portfolio, and Resume">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<meta http-equiv="Content-Language" content="en-US">%v
	</head>
	<body>
		<div id="wrap">
			<div class="container">
				<div class="header">
					<ul class="nav nav-pills pull-right">
						<li><a href="/">Home</a></li>
						<li><a href="/projects">Projects</a></li>
						<li><a href="/contact">Contact</a></li>
					</ul>
					<h3 class="text-muted"><a href="/">
						<img src="%%[1]v/img/lawas_logo_tight.png" width="122" height="50" alt="Lawas.Me" /></a>
					</h3>
				</div>
				%v
			</div>
		</div>
		<footer id="footer">
			<div class="container">
				<p class="credit"><a href="/">Home</a>&nbsp;&bull;&nbsp;<a href="/projects">Projects</a>&nbsp;&bull;&nbsp;<a href="/contact">Contact</a>&nbsp;&bull;&nbsp;<a href="/resume">Resume</a></p>
			</div>
		</footer>
		%v
		</body>
</html>`

var links string = `
		<link href="%[1]v/css/bootstrap/bootstrap.min.css" media="screen" rel="stylesheet" type="text/css">
		<link href="%[1]v/css/bootstrap/bootstrap-theme.min.css" media="screen" rel="stylesheet" type="text/css">
		<link href="%[1]v/css/style.css" media="screen" rel="stylesheet" type="text/css">
		<link href="%[1]v/img/favicon.ico" rel="shortcut icon" type="image/vnd.microsoft.icon">
		<link href="//stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet">
		`

var scripts string = `
	<!--[if lt IE 9]><script type="text/javascript" src="%[1]v/js/html5shiv.js"></script><![endif]-->
	<!--[if lt IE 9]><script type="text/javascript" src="%[1]v/js/respond.min.js"></script><![endif]-->
	<script type="text/javascript" src="%[1]v/js/jquery/jquery.min.js"></script>
	<script type="text/javascript" src="%[1]v/js/bootstrap/bootstrap.min.js"></script>
	<script type="text/javascript" src="%[1]v/js/gl/main.js"></script>
	<script>(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)})(window,document,'script','//www.google-analytics.com/analytics.js','ga');ga('create', 'UA-48615331-1', 'lawas.me');ga('send', 'pageview');</script>`

func indexAction(w http.ResponseWriter, r *http.Request) {
	body := fmt.Sprintf(document, links, `
			<div class="alpha-box">
				<div class="row">
					<div class="col-md-12">
						<div id="home-bio">
							<h3 style="margin:0 16px 0 0; width:160px; text-align: justify" class="pull-left"><span>Who Am I?</span><br><img src="%[1]v/img/gilberto-lawas-1.jpg" width="160" height="180" alt="Gilberto Lawas" class="img-thumbnail" style="margin-top: 10px;" ></h3>
							<p>In as few words as possible, I am a New York native living in Atlanta Georgia. Son to a single mother. Brother to four siblings. Father to five children. I spent the better part of a decade in service to our nation as a U.S. Marine before moving to Atlanta Georgia to take a job at a sewing machine parts distributor.</p>
							<p>The job was to make a web-site that integrated with the company inventory management system (MAS 90/200) to display current inventory information online <a href="http://web.archive.org/web/20071021032845/http://www.pd60.net/" target="_blank">www.pd60.net</a>. Which translated to roughly 80%% snapping inventory images, 10%% folding invoices, and 10%% composing HTML/JavaScript. But that 10%% was enough to get my foot in the door. </p>
							<p>That was about 14 years ago, and I have been building web applications ever since. First for a medical media company (<a href="http://catalog.nucleusinc.com/credits" target="_blank">Nucleus Medial Media</a>) as a PHP developer building / maintaining web applications. Later at a web design firm in Atlanta (<a href="http://www.airtightdesign.com" target="_blank">AirTight Design</a>) as a web developer; More recently as a DevOps engineer, at a global technology services company.</p>
							<p>It has been a long, successful and rewarding road. I have developed sites almost exclusively using the PHP language in a Linux, Apache, MySQL, PHP (LAMP) stack (though I've dabbled in JAVA, Python, and most recently GO Lang as needed).</p><p>Which brings me to the purpose of this site. The intent here is for this site is to serve as a kind of portfolio; To highlight some of the projects I have worked on; And to help me keep a record of my work. The &ldquo;projects&rdquo; portion will be a collection of links and dates with the occasional snippet where possible and appropriate.</p> 
							<div class="row">
								<blockquote class="col-md-8 col-md-offset-2" style="margin-top: 30px;"><a href="http://wiki.answers.com/Q/Origin_of_the_phrase_the_cobbler's_children_have_no_shoes" target="_BLANK"><font style="font-style: italic;font-size: 30px;color: darkslategray;">&ldquo;The cobbler&apos;s children have no shoes.&rdquo;</font></a>
									<p style="text-align: right;">- Source Unknown</p>
								</blockquote>
							</div>
							<div class="row">
								<p>That said, I would be remiss if I did not mention that this site is a personal project that will probably not get the attention it deserves. I will make an effort to, at the very minimum, make the site presentable. And then try to remind myself to keep it updated.</p>
							</div>
						</div>
					</div>
				</div>
			</div>`, scripts);
	fmt.Fprintf(w, body, cdn)
}

func projectsAction(w http.ResponseWriter, r *http.Request) {
	var scripts string = `
	<!--[if lt IE 9]><script type="text/javascript" src="%[1]v/js/html5shiv.js"></script><![endif]-->
	<!--[if lt IE 9]><script type="text/javascript" src="%[1]v/js/respond.min.js"></script><![endif]-->
	<script type="text/javascript" src="%[1]v/js/jquery/jquery.min.js"></script>
	<script type="text/javascript" src="%[1]v/js/bootstrap/bootstrap.min.js"></script>
	<script type="text/javascript" src="%[1]v/js/gl/projects.js"></script>
	<script>(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)})(window,document,'script','//www.google-analytics.com/analytics.js','ga');ga('create', 'UA-48615331-1', 'lawas.me');ga('send', 'pageview');</script>`

	body := fmt.Sprintf(document, links,`
				<div class="row">
					<div class="col-md-12 alpha-box" id="projects_wrapper" style="min-height: 600px">
						<div class="row">
							<div class="col-sm-8">
								<h3 class="first-title">What Have I Been Up To?</h3>
							</div>
							<div class="col-sm-4">
<!--								<div class="form-group" style="margin-top:16px;">
									<div class="input-group">
										<input class="form-control input-sm" type="text" placeholder="filter"></input>
										<div class="input-group-addon btn-info"><i class="glyphicon glyphicon-filter"></i></div>
									</div>
								</div> -->
							</div>
						</div>
						<hr style="margin-top:0; margin-bottom: 30px;">
					</div>
<!--					<div class="col-md-3">
						<ul>
							<li>Archive</li>
						</ul>
					</div> -->
				</div>`, scripts)

	fmt.Fprintf(w, body, cdn)
}

func contactAction(w http.ResponseWriter, r *http.Request) {
	body := fmt.Sprintf(document, links, `
				<div class="alpha-box">
					<div class='row'>
						<div class="col-md-4">
							<div class='gl-contact-form-wrapper'>
								<h3 class='first-title'>How To Get In Touch?<hr></h3>
								<div id="contact-social-links">
									<p>
										<a href="https://github.com/glawas" title="Gilberto Lawas GitHub Page" target="_BLANK"><b class="icon icon-github"></b>&nbsp;GitHub</a>
										<a href="//www.twitter.com/yahoodriver" title="Gilberto Lawas Twitter Page" target="_BLANK"><b class="icon icon-twitter"></b>&nbsp;Twitter</a><br>
										<a href="//bitbucket.org/glawas" title="Gilberto Lawas Bit Bucket Page" target="_BLANK"><b class="icon icon-bitbucket"></b>&nbsp;Bitbucket</a>
										<a href="//www.facebook.com/gilberto.lawas" title="Gilberto Lawas Facebook Page" target="_BLANK"><b class="icon icon-facebook"></b>&nbsp;Facebook</a>
									</p>
								</div>
								<div>
									<form method="POST" name="contact" id="contact-form" novalidate="novalidate" action="/contact">
										<div class="form-group">
											<input name="name" type="text" class="input-sm form-control" placeholder="Your Name" required="required" value="">
										</div>
										<div class="form-group">
											<input name="email" type="email" class="input-sm form-control" placeholder="Email Address" required="required" value="">
										</div>
										<div class="form-group">
											<textarea name="message" class="input-sm form-control" rows="8" placeholder="Your Message" required="required"></textarea>
										</div>
										<div class="form-group" class="checkbox" id="contact-form-action-wrapper">
											<input name="send" type="submit" id="contact-submit-button" class="btn btn-sm btn-default pull-right" value="Send">
											<input name="urgent_email" type="text" id="contact-urgent_email" value="">                
										</div>
									</form>
								</div>
							</div>
						</div>
						<div class='col-md-8'>	
							<div class="gl-contact-map-wrapper">
								<iframe width="100%" height="425" frameborder="0" scrolling="no" marginheight="0" marginwidth="0" src="https://maps.google.com/maps?q=conyers,+ga+30013&hl=en&ll=33.704706,-84.2,-84.235525&spn=0.062549,0.132093&sll=33.774828,-84.296312&sspn=0.062498,0.132093&t=m&hnear=Conyers,+Georgia+30013&z=10&output=embed"></iframe>
							</div>	
						</div>
					</div>
				</div>`, scripts);
	fmt.Fprintf(w, body, cdn)
}

func resumeAction(w http.ResponseWriter, r *http.Request) {
	body := fmt.Sprintf(document, links, `
				<div class="alpha-box">
					<div class="row">
						<div class="col-md-12">
							<div class='gl-resume-wrapper'>
								<h3 class="first-title">Where Have I Been?</h3>
								<ul class="nav nav-tabs">
									<li class="active"><a href="#resume-employment" data-toggle="tab">Employment</a></li>
									<li><a href="#resume-proficiencies" data-toggle="tab">Proficiencies</a></li>
									<li><a href="#resume-education" data-toggle="tab">Education</a></li>
									<li><a href="#resume-awards" data-toggle="tab">Awards</a></li>
								</ul>
								<div class="tab-content">
									<div id="resume-employment" class='tab-pane fade in active'>
										<h3>Relevant Experience</h3>
										<div class="row">
											<div class="col-md-8">
												<dl class='dl-horizontal'>
												<dt>2015 to Present</dt>
												<dd>Encompass Digital Media<br>DevOps Engineer<br>
													<ul>
														<li>Created automation scripts for server management and application deployments with PHP Deployer.</li>
														<li>Managed successful (re)launch of www.dvidshub.net, cms.dvidshub.net, api.dvidshub.net, and chat.dvidshub.net.</li>
														<li>Migrated to legacy server environments to lxc/d containers. </li>
													</ul>
												</dd>
												<dt>2013 to 2015</dt>
												<dd>Encompass Digital Media<br>Web Developer<br>
													<ul>
														<li>Updated CMS Authentication /Access Control Systems, to include the addition of two-factor authentication.</li>
														<li>Overhauled the application API and supporting documentation, to include the addition of OAuth2 authentication.</li>
														<li>Created WebRTC video chat application (JAVA).</li>
														<li>Overhauled various internal CMS applications.</li>
														<li>Created http/s image resizing utility (GO).</li>
													</ul>
												</dd>
												<dt>2011 to 2013</dt>
												<dd>AirTight Design, Atlanta, Georgia<br>Web Developer<br>
													<ul>
														<li>Developed custom web applications using PHP, JavaScript, HTML & CSS</li>
														<li>Deployed and managed application server environments.</li>
														<li>Maintained and updated existing web applications</li>
													</ul>
												</dd>
												<dt>2008 to 2011</dt>
												<dd>Nucleus Medical Media, Inc, Kennesaw, Georgia<br>Web Developer<br>
													<ul>
														<li>Developed dynamic web applications for both internal and external customers.</li>
														<li>Maintained and updated existing web applications.</li>
														<li>As part of a team, developed the catalog 2.0 framework.</li>
														<li>Assembled and managed the company Cinema 4D Net Render Farm.</li>
														<li>Provided workstation support as required.</li>
													</ul>
												</dd>
												<dt>2005 to 2007</dt>
												<dd>PD Sixty Distributors, Norcross, Georgia<br>Web Developer / Accounting Clerk / Sales Rep<br>
													<ul>
														<li>Designed, developed, maintained, and updated the company e-commerce website (www.pd60.net).</li>
														<li>Integrated company website with existing inventory & accounting software (MAS90/200).</li>
														<li>Photographed over 3000 inventory items to display on-line.</li>
														<li>Provided technical support for all web related inquires, including issuing and resetting log-on credentials.</li>
													</ul>
												</dd>
												<dt>1995 to 2004</dt>
												<dd>USMC MWSS-171, Iwakuni, Japan<br>Maintenance Administration Chief/Clerk<br>
													<ul>
														<li>Various management and training duties.</li>
														<li>Various clerical and administrative duties.</li>
														<li>Various physical duties.</li>
														<li>A complete duty list is available upon request.</li>
													</ul>
												</dd>
											</dl>
										</div>
									</div>
								</div>
								<div id="resume-proficiencies" class='tab-pane fade'>
									<h3>Technical Proficiencies</h3>
									<div class="row">
										<div class="col-md-7 col-md-offset-1">
											<h4>PHP</h4>
											<p>Over 10 years of procedural, object-oriented, and framework experience. Completed tasks ranging from simple Facebook page tabs to large-scale multi-node custom integrated web applications.</p>
											<h5>Frameworks</h5>
											<ul>
												<li>Zend Framework 2 (+6 months)</li>
												<li>Symfony 2 (+1 year)</li>
												<li>CakePHP  (+2 years)</li>
												<li>Zend Framework 1 (+5 years)</li>
											</ul>
										</div>
									</div>
									<div class="row">
										<div class="col-md-7 col-md-offset-1">
											<h4>Other Languages</h4>
											<ul>
												<li>JAVA (+6 months)</li>
												<li>Python (+6 months)</li>
												<li>GO  (+3 months)</li>
											</ul>
										</div>
									</div>
									<div class="row">
										<div class="col-md-7 col-md-offset-1">
											<h4>JavaScript</h4>
											<p>7 years of relative experience creating custom scripts with/without JS libraries.  Completed tasks ranging from simple mark-up manipulation and XHR submissions to complex browser based games and image manipulation.</p> 
											<h5>Libraries</h5>
											<ul>
												<li>JQuery (+6 years)</li>
												<li>DoJo (+3 months)</li>
											</ul>
										</div>
									</div>
									<div class="row">
										<div class="col-md-7 col-md-offset-1">
											<h4>Database</h4>
											<p>Over a decade of database experience; Most recently with MySql. Completed tasks ranging from simple data aggregation and custom views to stored procedures and complex reporting queries.</p>
											<p>I am not a DBA but I am very familiar with the planning, creation, optimization, and management of databases.</p>
											<ul>
												<li>MySql (+10 years)</li>
												<li>SOLR (+3 years)</li>
												<li>MSsql (+1 years - last used 2003 but it's like riding a bike)</li>
											</ul>
										</div>
									</div>
									<div class="row">
										<div class="col-md-7 col-md-offset-1">
											<h4>Server Administration</h4>
											<ul>
												<li>Ubuntu / Apache (+8 years)</li>
												<li>Centos / Apache (+2 years)</li>
												<li>PHP Deployer (+1 year)</li>
												<li>Docker, LXC container virtualization (+3 months)</li>
											</ul>
										</div>
									</div>
								</div>
								<div id="resume-education" class='tab-pane fade'>
						<h3>Education</h3>
						<div class="row">
							<div class="col-md-8">
								<dl class="dl-horizontal">
									<dt>2003</dt>
									<dd>University of Maryland University College Asia, Iwakuni, Japan<br>GPA 3.0</dd>
									<dt>2000</dt>
									<dd>Supervisory Maintenance Management, Okinawa, Japan<br>120 Hrs</dd>
									<dt>1998</dt>
									<dd>Leadership Course, Camp Lejeune, North Carolina<br>150 Hrs</dd>
									<dt>1995</dt>
									<dd>Basic Maintenance Management, Norfolk, Virginia<br>80 Hrs</dd>                
									<dt>1995</dt>
									<dd>Evander Childs High School, Bronx, New York<br>Graduated, General Studies</dt>
									<dt>1994</dt>
									<dd>B.O.C.E.S. Vo-Tech School, Goshen, NY<br>Graduated, Business Computer Technology</dt>
								</dl>
							</div>
						</div>
					<h3>Events</h3>
						<div class="row">
							<div class="col-md-8">
								<dl class="dl-horizontal">
									<dt>2016</dt>
									<dd>Participated in Hack-The-Pentagon (receiving end)</dt>
									<dd>Attended GreateWideOpen Atlanta - Tech Conference</dt>
								</dl>
							</div>
						</div>
					</div>
					<div id="resume-awards" class='tab-pane fade'>
						<h3>Awards</h3>
						<dl class='dl-horizontal'>
							<dt>1995-2001</dt>
							<dd>Marine Corps Good Conduct Medals (2)</dd>
		
							<dt>1999-2001</dt>
							<dd>Letter of Appreciation (3)</dd>
		
							<dt>1997</dt>
							<dd>Letter of Commendation</dd>
		
							<dt>1997</dt>
							<dd>Certificate of Commendation</dd>
		
							<dt>1995</dt>
							<dd>Meritorious Mast</dd>
						</dl>
					</div>
				</div>
				</div>
			</div>
		</div>
	</div>`, scripts);
	fmt.Fprintf(w, body, cdn)
}

func contactFormAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Conact Form Action");
}

func downloadAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DownloadAction");
}

func main() {

	http.HandleFunc("/projects", projectsAction)
	http.HandleFunc("/contact", contactAction)
	http.HandleFunc("/resume", resumeAction)
	http.HandleFunc("/download", downloadAction)
	http.HandleFunc("/", indexAction)

	env := os.Getenv("APP_ENV")
	
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		// cdn = ""
		http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./../public/css/"))))
		http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./../public/js"))))
		http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./../public/img"))))
		log.Println("Running api server in dev mode")
	}

	log.Fatal(http.ListenAndServe(":8686", nil))
}
