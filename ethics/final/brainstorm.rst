Brainstorming Using 9.3.1 Methodology
=====================================

Scenario
--------
You and other programmers are developing a communications system in disaster
zones (such as Manhattan on 9/11) for first responders. Users will be able to
communicate with each other, with supervisors near the scene, and with other
emergency personnel. The programmers plan to conduct the entirety of the
testing for the system in a field near the company office.

Decision to be made: To what extent do I stick to the initial testing plan?

Role of decision maker (me): A software developer

All the people & organizations affected (the stakeholders) + their rights
-------------------------------------------------------------------------
* First Responders, supervisors, and other emergency personnel
  * 4th amendments rights, ECPA rights

* Victims of a disaster
  * 4th amendments rights, ECPA rights

* My company
  * Myself (programmer)
  * Program Manager (the boss)
  * Coworkers

* Hackers

Risks, issues, problems, and consequences, and to whom they pertain
-------------------------------------------------------------------
System could fail at mission critical moments
* Cause response to be less organized
* Less efficient due to loss of communication

Privacy Violation
* People at the scene meant to help, could have their conversations monitored.

Hackers compromising the communications system
* these hackers could get caught
* hackers could put lives at risk if system is compromised during a disaster

The current testing plan
* Developers and Boss will produce lower quality work, overall less reliable


Most important - loss of life

Potential benefits. Identify who gets each benefit.
---------------------------------------------------
* Reliable communication
  - First responders will be able to do their job well

* Future of the company
  - Implementing good testing practices will improve the reliability of any
    software produced by my company

* Efficient first responders
  - More people are going to be helped, victims

* Team Competency
  - The testing strategy will improve the overall quality of work 



Cite applicable legal considerations and class terms
----------------------------------------------------
* Utilitarian reasoning
* High Reliability Organization
* Independent Verification and Validation
* ACM & IEEE
  - Software Engineering Code of Ethics and Professional Practice
* Decision Fatigue (A single person unit testing is hard.)
* 3 things to consider before leaking
* Computer Fraud and Abuse Act (CFAA) - 1984
* Key Aspects of Privacy
* Fair Information Principles



Identify vague or missing information in the scenario
-----------------------------------------------------
Who can access the communication lines? 
* My company and the government since the project recieves funding from the
  government. 

Is this platform being developed before or after the occurance of 9/11?
* For the purposes of this paper, I will say this initiative is happening after
  9/11.

What is the timeline for the development of this project?
* Understanding that a disaster can happen at any moment, for the purposes of
  this paper, I am going to assume that the developers have as much time as they
  need to complete the project.

Who is responsible for the decision to conduct the entirety of the testing for
the system in a field near the company office?
* For the purposes of this project, I'm going to say that it's the project
  manager is planning on testing the system in the field near the company
  office.

What is the willingness of the team (other programmers) and the project manager
to adapt?
* For the purposes of this project, the project manager truly believes the
  initial testing plan is sufficient, and will most likey, not budge. However,
  your coworkers, who are also professionals and have as much ownership of the
  project as you do, are more flexible when it comes to the development of the
  project. 

Analysis
========

1: Refuse to work on project unless testing procedure is improved, shows serious
 > Talk to my boss, the project manager explaining how as a computing
   professional, it is only right that the current testing policy is changed.
   > Yes: Overall code quality will improve
     > Work with my team to make a highly reliable system. This dynamic reflect
       highly reliable organization which is ideal for this system's particular
       use cases.
   > No: Consider leaking my company's testing plan. Forcus on 3 things to
     consider before leaking

2: Stick to current testing plan
 > Go through with the development of the platform and do entirety of testing in
   a field next to the office.
   > I could leak the state of the system. Similar to option 1. With the hopes
     of preventing its use in a disaster. 
   > Complete the project and remain silent of its shorcomings.

3: Ensure that the project is completed with adequate testing
 > Advocate for unit testing, ensuring that every line of code is tested.
   > Try to tell your coworks, fellow professionals, referencing the software
     ethic code that IEEE does. 
   > You get no help. You are ethically obligated to do as much testing as
     possible. You must put in the extra hours so the system is reliable and
     thoroughly tested. With utilitarian reasoning, any suffering you must go
     through to make this happen is worth it because the system will help people
     during a disaster. 

Synthesis
=========

Option 3 is the best. It's the one that give you the most valuable experience as
a developer and is best for all the people affected by the system you are
developing.
