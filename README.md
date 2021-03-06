DOT Problems
============

A simple web app to make reporting problems to Maine's DoT easier.

The Problem
-----------

The Maine DOT's website has a form to report a problem that is useful but not as user-friendly as it could be.  This is the form for my region: http://www.maine.gov/mdot/comments/rp/region4.htm  

As you can see, it simply says, "In the text box below, describe the problem you would like to report, and we will address it appropriately."  While this method certainly works, it would be far easier to incorporate geospatial features such as a map or even a mobile app that will upload geotags as to the location of the problem.  

Perhaps resources such as openstreetmaps.org, ESRI, or others can be of use. 

The Solution
-------------

The solution app should:

  1. Provide geo-location look up when the page/app is started
  2. Duplicate the current form on the Maine.gov reporting page
  3. Ask for your email address and comment
  4. Infer the zipcode + date
  5. Dispatch the form data to the Maine DoT
  6. Marvel at how much easier that was
 
Thankfully, Maine DoT doesn't worry about cross-site scripting attacks, so I
was able to hijack their form and post to it from my HTML page. 
