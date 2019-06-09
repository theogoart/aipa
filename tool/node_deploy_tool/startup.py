#coding=utf-8

import sys
import os, time, stat
import types
import exceptions, traceback
import shutil

import urllib
import urllib2

import tarfile

from subprocess import Popen, PIPE, STDOUT

#############################################################
##             USER CONFIGURATIONS                          #
#############################################################
# aipa code dir                                           # 
GLOBAL_aipa_DIR = '/home/aipa/aipa_dir'               #
aipa_PROGRAM_WORK_DIR   = GLOBAL_aipa_DIR + '/work_dir' #
#sequences' change is not allowed                           #         
user_choice_list = {   'install_base' : 'yes',              #
		       'install_golang': 'yes',             #
                       'install_mongodb':'no',		    #
                       'install_gomicro': 'no',             #
                       'install_aipa_source_code': 'no'   #
                   }	                                    #
							    #
GOPATH = '/home/aipa/go'                                  #
GOROOT = '/usr/lib/go'					    #	
#############################################################

def get_all_install_options():
    for item in option_list:
        yield item

def predo_cmd(cmd, *optional):
    stderr = ''
    print cmd
    process = Popen(cmd, stdout=PIPE, stderr=PIPE, shell=True)

    if not 'no_wait' in optional:
            while Popen.poll(process) == None:
                r = process.stdout.readline().strip().decode('utf-8')
                if r:
                    print(r);
                    print(process.stdout.readline().strip().decode('utf-8'))
            _, stderr = process.communicate()


if __name__ == '__main__':

    _ = GOPATH
    _ = GLOBAL_aipa_DIR
    _ = aipa_PROGRAM_WORK_DIR		

    if len(sys.argv) <= 1 or sys.argv[1] in ['--help', '-h']:
	   common.print_help()
	   exit(1)
    
    conf = '.installation_config.txt'
    if os.path.exists(conf):
		with open(conf, 'r') as f:
			for line in f.readlines():
				if 'GLOBAL_aipa_DIR' in line.split(':\n'):
					GLOBAL_aipa_DIR = line.split(':\n')[1]
				elif 'aipa_PROGRAM_WORK_DIR' in line.split(':\n'):
					aipa_PROGRAM_WORK_DIR = line.split(':\n')[1]
    				
    elif not sys.argv[1] == 'install':
	print '\n****** WARNING: you have no .installation_config.txt file under current directory, so use default GLOBAL_aipa_DIR and aipa_PROGRAM_WORK_DIR ***********\n'    
		
    if sys.argv[1] == 'install':
	   if os.geteuid() != 0:
                       print "The required packages must be installed under root user. Please turn into root account first."
                       exit(1)
	   x=raw_input('Do you need to install node by recommanded configuration? Y/N')
	   if x.upper() in ('Y', 'YES'):
		pass
	   elif x.upper() in ('N', 'NO'):
                
		x=raw_input('Please choose your working directory: default is %s ' %GLOBAL_aipa_DIR )
		if len(x) > 0:   
			GLOBAL_aipa_DIR = x
		x=raw_input('Please choose your aipa programe files\' directory: default is %s ' % aipa_PROGRAM_WORK_DIR )   
		if len(x) > 0:
			aipa_PROGRAM_WORK_DIR = x
		   
		if not os.path.exists(aipa_PROGRAM_WORK_DIR):
			os.makedirs(aipa_PROGRAM_WORK_DIR)
		   
		x=raw_input('Do you need to install base packages? Y/N: default is %s ' % user_choice_list['install_base']) 
		if len(x) > 0:
			if x.upper() in ('Y', 'YES'):   
				user_choice_list['install_base'] = 'yes'
			elif x.upper() in ('N', 'NO'):
				user_choice_list['install_base'] = 'no'
		else:
			pass
		   
	   else:
		print 'Wrong input. Please try again.'
		exit(1)
	   		   
	   	
	   print 'Your install configuration is as following: \n'	   	
	   for key, item in user_choice_list.items():
		if item is 'yes':
			print key
			print
	   	
	   print 'Your working directory: %s' %GLOBAL_aipa_DIR
           print 'Your aipa programe files\' directory: %s' % aipa_PROGRAM_WORK_DIR
	   print
	   
           x=raw_input('Are you sure? Y/N')
	   if  x.upper() in ('Y', 'YES'):
		
	    	print '\n============ NOW INSTALL aipa ENVIRONMENT =======================\n'
	   	deploy_obj = aipa_node_deploy()
	   	deploy_obj.install_env()
	   else:
		print 'All right. Bye bye.'
		exit(1)	
    else:
    	#obj = download_progress_bar('https://studygolang.com/dl/golang/go1.10.1.linux-amd64.tar.gz') 
    	#obj.download_with_progressbar('./go1.10.1.linux-amd64.tar.gz')
	common.print_help()

    exit(0)


