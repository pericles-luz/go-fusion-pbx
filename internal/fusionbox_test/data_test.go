package fusionbox_test

func callcenterWithBranchPageContent() string {
	return `
	
	<!DOCTYPE html>
	
	<html xmlns='http://www.w3.org/1999/xhtml' xml:lang='en' lang='en'>
	<head>
	<meta charset='utf-8'>
	<meta http-equiv='Content-Type' content='text/html; charset=UTF-8'>
	<meta http-equiv='X-UA-Compatible' content='IE=edge'>
	<meta name='viewport' content='width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no' />
	<meta name="robots" content="noindex, nofollow, noarchive" />
	
		<link rel='stylesheet' type='text/css' href='/resources/bootstrap/css/bootstrap.min.css.php'>
		<link rel='stylesheet' type='text/css' href='/resources/bootstrap/css/bootstrap-tempusdominus.min.css.php'>
		<link rel='stylesheet' type='text/css' href='/resources/bootstrap/css/bootstrap-colorpicker.min.css.php'>
		<link rel='stylesheet' type='text/css' href='/resources/fontawesome/css/all.min.css.php'>
		<link rel='stylesheet' type='text/css' href='/themes/default/css.php'>
		
		<link rel='icon' href='/themes/default/favicon.ico'>
	
		<title>Call Center Queue - FusionPBX</title>
	
		<script language='JavaScript' type='text/javascript' src='/resources/jquery/jquery.min.js.php'></script>
		<script language='JavaScript' type='text/javascript' src='/resources/jquery/jquery.autosize.input.js.php'></script>
		<script language='JavaScript' type='text/javascript' src='/resources/momentjs/moment-with-locales.min.js.php'></script>
		<script language='JavaScript' type='text/javascript' src='/resources/bootstrap/js/bootstrap.min.js.php'></script>
		<script language='JavaScript' type='text/javascript' src='/resources/bootstrap/js/bootstrap-tempusdominus.min.js.php'></script>
		<script language='JavaScript' type='text/javascript' src='/resources/bootstrap/js/bootstrap-colorpicker.min.js.php'></script>
		<script language='JavaScript' type='text/javascript' src='/resources/bootstrap/js/bootstrap-pwstrength.min.js.php'></script>
		<script language='JavaScript' type='text/javascript'>window.FontAwesomeConfig = { autoReplaceSvg: false }</script>
		<script language='JavaScript' type='text/javascript' src='/resources/fontawesome/js/solid.min.js.php' defer></script>
	
		
		<script language='JavaScript' type='text/javascript'>
	
			//message bar display
				
				function display_message(msg, mood, delay) {
					mood = mood !== undefined ? mood : 'default';
					delay = delay !== undefined ? delay : 1000;
					if (msg !== '') {
						var message_text = $(document.createElement('div'));
						message_text.addClass('message_text message_mood_'+mood);
						message_text.html(msg);
						message_text.on('click', function() {
							var object = $(this);
							object.clearQueue().finish();
							$('#message_container div').remove();
							$('#message_container').css({opacity: 0, 'height': 0}).css({'height': 'auto'});
						} );
						$('#message_container').append(message_text);
						message_text.css({'height': 'auto'}).animate({opacity: 1}, 250, function(){
							$('#message_container').delay(delay).animate({opacity: 0, 'height': 0}, 500, function() {
								$('#message_container div').remove();
								$('#message_container').animate({opacity: 1}, 300).css({'height': 'auto'});
							});
						});
					}
				}
				
	
			
		
		$(document).ready(function() {
			
	
					//render the messages
	
	
			//message bar hide on hover
				
				$('#message_container').on('mouseenter',function() {
					$('#message_container div').remove();
					$('#message_container').css({opacity: 0, 'height': 0}).css({'height': 'auto'});
				});
				
	
			//domain selector controls
								
					$('.domain_selector_domain').on('click', function() { show_domains(); });
					$('#header_domain_selector_domain').on('click', function() { show_domains(); });
					$('#domains_hide').on('click', function() { hide_domains(); });
	
					function show_domains() {
						search_domains('domains_list');
	
						$('#domains_visible').val(1);
						var scrollbar_width = (window.innerWidth - $(window).width()); //gold: only solution that worked with body { overflow:auto } (add -ms-overflow-style: scrollbar; to <body> style for ie 10+)
						if (scrollbar_width > 0) {
							$('body').css({'margin-right':scrollbar_width, 'overflow':'hidden'}); //disable body scroll bars
							$('.navbar').css('margin-right',scrollbar_width); //adjust navbar margin to compensate
							$('#domains_container').css('right',-scrollbar_width); //domain container right position to compensate
						}
						$(document).scrollTop(0);
						$('#domains_container').show();
						$('#domains_block').animate({marginRight: '+=300'}, 400, function() {
							$('#domains_search').trigger('focus');
						});
					}
	
					function hide_domains() {
						$('#domains_visible').val(0);
						$(document).ready(function() {
							$('#domains_block').animate({marginRight: '-=300'}, 400, function() {
								$('#domains_search').val('');
								$('.navbar').css('margin-right','0'); //restore navbar margin
								$('#domains_container').css('right','0'); //domain container right position
								$('#domains_container').hide();
								$('body').css({'margin-right':'0','overflow':'auto'}); //enable body scroll bars
								document.activeElement.blur();
							});
						});
					}
					
				
			//keyboard shortcut scripts
	
			//key: [enter] - retain default behavior to submit form, when present - note: safari does not honor the first submit element when hiding it using 'display: none;' in the setAttribute method
								
					var action_bar_actions, first_form, first_submit, modal_input_class, modal_continue_button;
					action_bar_actions = document.querySelector('div#action_bar.action_bar > div.actions');
					first_form = document.querySelector('form#frm');
	
					if (action_bar_actions !== null) {
						if (first_form !== null) {
							first_submit = document.createElement('input');
							first_submit.type = 'submit';
							first_submit.id = 'default_submit';
							first_submit.setAttribute('style','position: absolute; left: -10000px; top: auto; width: 1px; height: 1px; overflow: hidden;');
							first_form.prepend(first_submit);
							window.addEventListener('keydown',function(e){
								modal_input_class = e.target.className;
								if (e.which == 13 && (e.target.tagName == 'INPUT' || e.target.tagName == 'SELECT')) {
									if (modal_input_class.includes('modal-input')) {
										e.preventDefault();
										modal_continue_button = document.getElementById(e.target.dataset.continue);
										if (modal_continue_button) { modal_continue_button.click(); }
									}
									else {
										if (typeof window.submit_form === 'function') { submit_form(); }
										else { document.getElementById('frm').submit(); }
									}
								}
							});
						}
					}
					
				
			//common (used by delete and toggle)
								var list_checkboxes;
					list_checkboxes = document.querySelectorAll('table.list tr.list-row td.checkbox input[type=checkbox]');
				
			//keyup event listener
				
				window.addEventListener('keyup', function(e) {
					
	
					//key: [escape] - close modal window, if open, or toggle domain selector
						
						if (e.which == 27) {
							e.preventDefault();
							var modals, modal_visible, modal;
							modal_visible = false;
							modals = document.querySelectorAll('div.modal-window');
							if (modals.length !== 0) {
								for (var x = 0, max = modals.length; x < max; x++) {
									modal = document.getElementById(modals[x].id);
									if (window.getComputedStyle(modal).getPropertyValue('opacity') == 1) {
										modal_visible = true;
									}
								}
							}
							if (modal_visible) {
								modal_close();
							}
							
														
								else {
									if (document.getElementById('domains_visible').value == 0) {
										show_domains();
									}
									else {
										hide_domains();
									}
								}
								
													
						}
						
	
					//key: [insert], list: to add
												
							if (e.which == 45 && !(e.target.tagName == 'INPUT' && e.target.type == 'text') && e.target.tagName != 'TEXTAREA') {
								e.preventDefault();
								var add_button;
								add_button = document.getElementById('btn_add');
								if (add_button === null || add_button === undefined) {
									add_button = document.querySelector('button[name=btn_add]');
								}
								if (add_button !== null) { add_button.click(); }
							}
							
						
					//key: [delete], list: to delete checked, edit: to delete
						
					//key: [space], list,edit:prevent default space key behavior when opening toggle confirmation (which would automatically *click* the focused continue button on key-up)
												
							if (e.which == 32 && e.target.id == 'btn_toggle') {
								e.preventDefault();
							}
							
						
			//keyup end
				
				});
				
	
			//keydown event listener
				
				window.addEventListener('keydown', function(e) {
					
	
					//key: [space], list: to toggle checked - note: for default [space] checkbox behavior (ie. toggle focused checkbox) include in the if statement: && !(e.target.tagName == 'INPUT' && e.target.type == 'checkbox')
												
							if (e.which == 32 && !(e.target.tagName == 'INPUT' && e.target.type == 'text') && e.target.tagName != 'BUTTON' && !(e.target.tagName == 'INPUT' && e.target.type == 'button') && !(e.target.tagName == 'INPUT' && e.target.type == 'submit') && e.target.tagName != 'TEXTAREA' && list_checkboxes.length !== 0) {
								e.preventDefault();
								var toggle_button;
								toggle_button = document.querySelector('button[name=btn_toggle]');
								if (toggle_button === null || toggle_button === undefined) {
									toggle_button = document.getElementById('btn_toggle');
								}
								if (toggle_button !== null) { toggle_button.click(); }
							}
							
						
					//key: [ctrl]+[a], list,edit: to check all
												
							if ((((e.which == 97 || e.which == 65) && (e.ctrlKey || e.metaKey) && !e.shiftKey) || e.which == 19) && !(e.target.tagName == 'INPUT' && e.target.type == 'text') && e.target.tagName != 'TEXTAREA') {
								var all_checkboxes;
								all_checkboxes = document.querySelectorAll('table.list tr.list-header th.checkbox input[name=checkbox_all]');
								if (typeof all_checkboxes != 'object' || all_checkboxes.length == 0) {
									all_checkboxes = document.querySelectorAll('td.edit_delete_checkbox_all > span > input[name=checkbox_all]');
								}
								if (typeof all_checkboxes == 'object' && all_checkboxes.length > 0) {
									e.preventDefault();
									for (var x = 0, max = all_checkboxes.length; x < max; x++) {
										all_checkboxes[x].click();
									}
								}
							}
							
	
						
					//key: [ctrl]+[s], edit: to save
												
							if (((e.which == 115 || e.which == 83) && (e.ctrlKey || e.metaKey) && !e.shiftKey) || (e.which == 19)) {
								e.preventDefault();
								var save_button;
								save_button = document.getElementById('btn_save');
								if (save_button === null || save_button === undefined) {
									save_button = document.querySelector('button[name=btn_save]');
								}
								if (save_button !== null) { save_button.click(); }
							}
							
						
					//key: [ctrl]+[c], list,edit: to copy
																			
								if (
									(
										(
											(e.which == 99 || e.which == 67) &&
											(e.ctrlKey || e.metaKey) &&
											!e.shiftKey
										) ||
										e.which == 19
									) &&
									!(e.target.tagName == 'INPUT' && e.target.type == 'text') &&
									e.target.tagName != 'TEXTAREA'
									) {
								
													
								var current_selection, copy_button;
								current_selection = window.getSelection();
								if (current_selection === null || current_selection === undefined || current_selection.toString() == '') {
									e.preventDefault();
									copy_button = document.querySelector('button[name=btn_copy]');
									if (copy_button === null || copy_button === undefined) {
										copy_button = document.getElementById('btn_copy');
									}
									if (copy_button !== null) { copy_button.click(); }
								}
							}
							
						
			//keydown end
				
				});
				
	
			//link list rows
				
				$('.tr_hover tr,.list tr').each(function(i,e) {
					$(e).children('td:not(.list_control_icon,.list_control_icons,.tr_link_void,.list-row > .no-link,.list-row > .checkbox,.list-row > .button,.list-row > .action-button)').on('click', function() {
						var href = $(this).closest('tr').attr('href');
						var target = $(this).closest('tr').attr('target');
						if (href) {
							if (target) { window.open(href, target); }
							else { window.location = href; }
						}
					});
				});
				
	
			//autosize jquery autosize plugin on applicable input fields
				
				$('input[type=text].txt.auto-size,input[type=number].txt.auto-size,input[type=password].txt.auto-size,input[type=text].formfld.auto-size,input[type=number].formfld.auto-size,input[type=password].formfld.auto-size').autosizeInput();
				
	
			//initialize bootstrap tempusdominus (calendar/datetime picker) plugin
				
				$(function() {
					//set defaults
						$.fn.datetimepicker.Constructor.Default = $.extend({}, $.fn.datetimepicker.Constructor.Default, {
							buttons: {
								showToday: true,
								showClear: true,
								showClose: true,
							},
							icons: {
								time: 'fas fa-clock',
								date: 'fas fa-calendar-alt',
								up: 'fas fa-arrow-up',
								down: 'fas fa-arrow-down',
								previous: 'fas fa-chevron-left',
								next: 'fas fa-chevron-right',
								today: 'fas fa-calendar-check',
								clear: 'fas fa-trash',
								close: 'fas fa-times',
							}
						});
					//define formatting of individual classes
						$('.datepicker').datetimepicker({ format: 'YYYY-MM-DD', });
						$('.datetimepicker').datetimepicker({ format: 'YYYY-MM-DD HH:mm', });
						$('.datetimepicker-future').datetimepicker({ format: 'YYYY-MM-DD HH:mm', minDate: new Date(), });
						$('.datetimesecpicker').datetimepicker({ format: 'YYYY-MM-DD HH:mm:ss', });
				});
				
	
			//apply bootstrap colorpicker plugin
				
				$(function(){
					$('.colorpicker').colorpicker({
						align: 'left',
						customClass: 'colorpicker-2x',
						sliders: {
							saturation: {
								maxLeft: 200,
								maxTop: 200
							},
							hue: {
								maxTop: 200
							},
							alpha: {
								maxTop: 200
							}
						}
					});
				});
				
	
			//apply bootstrap password strength plugin
				
				$('#password').pwstrength({
					common: {
						minChar: 8,
						usernameField: '#username',
					},
					//rules: { },
					ui: {
						colorClasses: ['danger', 'warning', 'warning', 'warning', 'success', 'success'], //weak,poor,normal,medium,good,strong
						progressBarMinPercentage: 15,
						showVerdicts: false,
						viewports: {
							progress: '#pwstrength_progress'
						}
					}
				});
				
	
			//crossfade menu brand images (if hover version set)
				
			//generate resizeEnd event after window resize event finishes (used when side menu and on messages app)
				
				$(window).on('resize', function() {
					if (this.resizeTO) { clearTimeout(this.resizeTO); }
					this.resizeTO = setTimeout(function() { $(this).trigger('resizeEnd'); }, 180);
				});
				
	
			//side menu: adjust content container width after window resize
				
		
		}); //document ready end
		
	
	
		//audio playback functions
			
			var recording_audio, audio_clock;
	
			function recording_play(recording_id) {
				if (document.getElementById('recording_progress_bar_'+recording_id)) {
					document.getElementById('recording_progress_bar_'+recording_id).style.display='';
				}
				recording_audio = document.getElementById('recording_audio_'+recording_id);
	
				if (recording_audio.paused) {
					recording_audio.volume = 1;
					recording_audio.play();
					document.getElementById('recording_button_'+recording_id).innerHTML = "<span class='fas fa-pause fa-fw'></span>";
					audio_clock = setInterval(function () { update_progress(recording_id); }, 20);
	
					$('[id*=recording_button]').not('[id*=recording_button_'+recording_id+']').html("<span class='fas fa-play fa-fw'></span>");
					$('[id*=recording_progress_bar]').not('[id*=recording_progress_bar_'+recording_id+']').css('display', 'none');
	
					$('audio').each(function(){$('#menu_side_container').width()
						if ($(this).get(0) != recording_audio) {
							$(this).get(0).pause(); //stop playing
							$(this).get(0).currentTime = 0; //reset time
						}
					});
				}
				else {
					recording_audio.pause();
					document.getElementById('recording_button_'+recording_id).innerHTML = "<span class='fas fa-play fa-fw'></span>";
					clearInterval(audio_clock);
				}
			}
	
			function recording_stop(recording_id) {
				recording_reset(recording_id);
				clearInterval(audio_clock);
			}
	
			function recording_reset(recording_id) {
				recording_audio = document.getElementById('recording_audio_'+recording_id);
				recording_audio.pause();
				recording_audio.currentTime = 0;
				if (document.getElementById('recording_progress_bar_'+recording_id)) {
					document.getElementById('recording_progress_bar_'+recording_id).style.display='none';
				}
				document.getElementById('recording_button_'+recording_id).innerHTML = "<span class='fas fa-play fa-fw'></span>";
				clearInterval(audio_clock);
			}
	
			function update_progress(recording_id) {
				recording_audio = document.getElementById('recording_audio_'+recording_id);
				var recording_progress = document.getElementById('recording_progress_'+recording_id);
				var value = 0;
				if (recording_audio.currentTime > 0) {
					value = (100 / recording_audio.duration) * recording_audio.currentTime;
				}
				recording_progress.style.marginLeft = value + '%';
				if (parseInt(recording_audio.duration) > 30) { //seconds
					clearInterval(audio_clock);
				}
			}
			
	
		//handle action bar style on scroll
			
			window.addEventListener('scroll', function(){
				action_bar_scroll('action_bar', 20);
			}, false);
			function action_bar_scroll(action_bar_id, scroll_position, function_sticky, function_inline) {
				if (document.getElementById(action_bar_id)) {
					//sticky
						if (this.scrollY > scroll_position) {
							document.getElementById(action_bar_id).classList.add('scroll');
							if (typeof function_sticky === 'function') { function_sticky(); }
						}
					//inline
						if (this.scrollY < scroll_position) {
							document.getElementById(action_bar_id).classList.remove('scroll');
							if (typeof function_inline === 'function') { function_inline(); }
						}
				}
			}
			
	
		//enable button class button
			
			function button_enable(button_id) {
				button = document.getElementById(button_id);
				button.disabled = false;
				button.classList.remove('disabled');
				if (button.parentElement.nodeName == 'A') {
					anchor = button.parentElement;
					anchor.classList.remove('disabled');
					anchor.setAttribute('onclick','');
				}
			}
			
	
		//disable button class button
			
			function button_disable(button_id) {
				button = document.getElementById(button_id);
				button.disabled = true;
				button.classList.add('disabled');
				if (button.parentElement.nodeName == 'A') {
					anchor = button.parentElement;
					anchor.classList.add('disabled');
					anchor.setAttribute('onclick','return false;');
				}
			}
			
	
		//checkbox on change
			
			function checkbox_on_change(checkbox) {
				checked = false;
				var inputs = document.getElementsByTagName('input');
				for (var i = 0, max = inputs.length; i < max; i++) {
					if (inputs[i].type === 'checkbox' && inputs[i].checked == true) {
						checked = true;
						break;
					}
				}
				btn_copy = document.getElementById("btn_copy");
				btn_toggle = document.getElementById("btn_toggle");
				btn_delete = document.getElementById("btn_delete");
				if (checked == true) {
					if (btn_copy) {
						btn_copy.style.display = "inline";
					}
					if (btn_toggle) {
						btn_toggle.style.display = "inline";
					}
					if (btn_delete) {
						btn_delete.style.display = "inline";
					}
				}
				 else {
					if (btn_copy) {
						btn_copy.style.display = "none";
					}
					if (btn_toggle) {
						btn_toggle.style.display = "none";
					}
					if (btn_delete) {
						btn_delete.style.display = "none";
					}
				 }
			}
			
	
		//list page functions
			
			function list_all_toggle(modifier) {
				var checkboxes = (modifier !== undefined) ? document.getElementsByClassName('checkbox_'+modifier) : document.querySelectorAll("input[type='checkbox']");
				var checkbox_checked = document.getElementById('checkbox_all' + (modifier !== undefined ? '_'+modifier : '')).checked;
				for (var i = 0, max = checkboxes.length; i < max; i++) {
					checkboxes[i].checked = checkbox_checked;
				}
				if (document.getElementById('btn_check_all') && document.getElementById('btn_check_none')) {
					if (checkbox_checked) {
						document.getElementById('btn_check_all').style.display = 'none';
						document.getElementById('btn_check_none').style.display = '';
					}
					else {
						document.getElementById('btn_check_all').style.display = '';
						document.getElementById('btn_check_none').style.display = 'none';
					}
				}
			}
	
			function list_all_check() {
				var inputs = document.getElementsByTagName('input');
				document.getElementById('checkbox_all').checked;
				for (var i = 0, max = inputs.length; i < max; i++) {
					if (inputs[i].type === 'checkbox') {
						inputs[i].checked = true;
					}
				}
			}
	
			function list_self_check(checkbox_id) {
				var inputs = document.getElementsByTagName('input');
				for (var i = 0, max = inputs.length; i < max; i++) {
					if (inputs[i].type === 'checkbox' && inputs[i].name.search['enabled'] == -1) {
						inputs[i].checked = false;
					}
				}
				document.getElementById(checkbox_id).checked = true;
			}
	
			function list_action_set(action) {
				document.getElementById('action').value = action;
			}
	
			function list_form_submit(form_id) {
				document.getElementById(form_id).submit();
			}
	
			function list_search_reset() {
				document.getElementById('btn_reset').style.display = 'none';
				document.getElementById('btn_search').style.display = '';
			}
			
	
		//edit page functions
			
			function edit_all_toggle(modifier) {
				var checkboxes = document.getElementsByClassName('checkbox_'+modifier);
				var checkbox_checked = document.getElementById('checkbox_all_'+modifier).checked;
				if (checkboxes.length > 0) {
					for (var i = 0; i < checkboxes.length; ++i) {
						checkboxes[i].checked = checkbox_checked;
					}
					if (document.getElementById('btn_delete')) {
						document.getElementById('btn_delete').value = checkbox_checked ? '' : 'delete';
					}
				}
			}
	
			function edit_delete_action(modifier) {
				var checkboxes = document.getElementsByClassName('chk_delete');
				if (document.getElementById('btn_delete') && checkboxes.length > 0) {
					var checkbox_checked = false;
					for (var i = 0; i < checkboxes.length; ++i) {
						if (checkboxes[i].checked) {
							checkbox_checked = true;
						}
						else {
							if (document.getElementById('checkbox_all'+(modifier !== undefined ? '_'+modifier : ''))) {
								document.getElementById('checkbox_all'+(modifier !== undefined ? '_'+modifier : '')).checked = false;
							}
						}
					}
					document.getElementById('btn_delete').value = checkbox_checked ? '' : 'delete';
				}
			}
			
	
		//modal functions
			
			function modal_open(modal_id, focus_id) {
				var modal = document.getElementById(modal_id);
				modal.style.opacity = '1';
				modal.style.pointerEvents = 'auto';
				if (focus_id !== undefined) {
					document.getElementById(focus_id).focus();
				}
			}
	
			function modal_close() {
				var modals = document.getElementsByClassName('modal-window');
				if (modals.length > 0) {
					for (var m = 0; m < modals.length; ++m) {
						modals[m].style.opacity = '0';
						modals[m].style.pointerEvents = 'none';
					}
				}
				document.activeElement.blur();
			}
			
	
		//misc functions
			
			function swap_display(a_id, b_id, display_value) {
				display_value = display_value !== undefined ? display_value : 'inline-block';
				a = document.getElementById(a_id);
				b = document.getElementById(b_id);
				if (window.getComputedStyle(a).display === 'none') {
					a.style.display = display_value;
					b.style.display = 'none';
				}
				else {
					a.style.display = 'none';
					b.style.display = display_value;
				}
			}
	
			function hide_password_fields() {
				var password_fields = document.querySelectorAll("input[type='password']");
				for (var p = 0, max = password_fields.length; p < max; p++) {
					password_fields[p].style.visibility = 'hidden';
					password_fields[p].type = 'text';
				}
			}
	
			window.addEventListener('beforeunload', function(e){
				hide_password_fields();
			});
			
	
				
			function search_domains(element_id) {
			var xhttp = new XMLHttpRequest();
			xhttp.onreadystatechange = function() {
				//if (this.readyState == 4 && this.status == 200) {
				//	document.getElementById(element_id).innerHTML = this.responseText;
				//}
	
				//remove current options
				document.getElementById(element_id).innerHTML = '';
	
				if (this.readyState == 4 && this.status == 200) {
	
					//create the json object from the response
					obj = JSON.parse(this.responseText);
	
					//update the domain count
					document.getElementById('domain_count').innerText = '('+ obj.length +')';
	
					//add new options from the json results
					for (var i=0; i < obj.length; i++) {
						
						//get the variables
						domain_uuid = obj[i].domain_uuid;
						domain_name = obj[i].domain_name;
						if (obj[i].domain_description != null) {
						//	domain_description = DOMPurify.sanitize(obj[i].domain_description);
						}
	
						//create a div element
						var div = document.createElement('div');
	
						//add a div title
						div.title = obj[i].domain_name;
	
						//add a css class
						div.classList.add("domains_list_item");
	
						//alternate the background color
						if(i%2==0) {
							div.style.background = '#eaedf2';
						}
						else {
							div.style.background = '#ffffff';
						}
	
						//set the active domain style 
						if ('46ce3c4e-4d9e-420b-9190-116820650089' == obj[i].domain_uuid) {
							div.style.background = '#eeffee';
							div.style.fontWeight = 'bold';
							//div.classList.add("domains_list_item_active");
							//var item_description_class = 'domain_active_list_item_description';
						}
						else {
							//div.classList.add("domains_list_item_inactive");
							//var item_description_class = 'domain_inactive_list_item_description';
						}
	
						//set link on domain div in list
						div.setAttribute('onclick',"window.location.href='/core/domains/domains.php?domain_uuid=" + obj[i].domain_uuid + "&domain_change=true';");
	
						//define domain link text and description (if any)
						link_label = obj[i].domain_name;
						if (obj[i].domain_description != null) {
							link_label += " <span class='domain_list_item_description' title=\"" + obj[i].domain_description + "\">" + obj[i].domain_description + "</span>";
						}
						var a_tag = document.createElement('a');
						a_tag.setAttribute('href','manage:'+obj[i].domain_name);
						a_tag.setAttribute('onclick','event.preventDefault();');
						a_tag.innerHTML = link_label;
						div.appendChild(a_tag);
	
						document.getElementById(element_id).appendChild(div);
					}
				}
			};
			search = document.getElementById('domains_search');
			if (search.value) {
				//xhttp.open("GET", "/core/domains/domain_list.php?search="+search.value, true);
				xhttp.open("GET", "/core/domains/domain_json.php?search="+search.value+"&9424be1701c50fa8a587afd380635e8cd89c9dccbb70f7f0b066bab5b1687938=16d90aa7da0fab59b481135faf22cbc88c6fe8c2d354e7ae2be9e7b763a29c04", true);
			}
			else {
				//xhttp.open("GET", "/core/domains/domain_list.php", true);
				xhttp.open("GET", "/core/domains/domain_json.php?9424be1701c50fa8a587afd380635e8cd89c9dccbb70f7f0b066bab5b1687938=16d90aa7da0fab59b481135faf22cbc88c6fe8c2d354e7ae2be9e7b763a29c04", true);
			}
			xhttp.send();
		}
			</script>
	
	</head>
	<body>
	
				<div id='message_container'></div>
	
				
				<div id='domains_container'>
					<input type='hidden' id='domains_visible' value='0'>
					<div id='domains_block'>
						<div id='domains_header'>
							<input id='domains_hide' type='button' class='btn' style='float: right' value="Close">
							<a id='domains_title' href='/core/domains/domains.php'>Domains <span id='domain_count' style='font-size: 80%;'></span></a>
							<br><br>
							<input type='text' id='domains_search' class='formfld' style='margin-left: 0; min-width: 100%; width: 100%;' placeholder="Search..." onkeyup="search_domains('domains_list');">
						</div>
						<div id='domains_list'></div>
					</div>
				</div>
	
			
				<div id='qr_code_container' style='display: none;' onclick='$(this).fadeOut(400);'>
				<table cellpadding='0' cellspacing='0' border='0' width='100%' height='100%'><tr><td align='center' valign='middle'>
					<span id='qr_code' onclick="$('#qr_code_container').fadeOut(400);"></span>
				</td></tr></table>
			</div>
	
											 <nav class='navbar navbar-expand-sm fixed-top' >
		<div class='container-fluid' style='width: calc(90% - 20px); padding: 0;'>
			<div class='navbar-brand'>
				<a href='/'>				<img id='menu_brand_image' class='navbar-logo' src='/themes/default/images/logo.png' title="FusionPBX"></a>
				<a style='margin: 0;'></a>
			</div>
			<button type='button' class='navbar-toggler' data-toggle='collapse' data-target='#main_navbar' aria-expanded='false' aria-controls='main_navbar' aria-label='Toggle Menu'>
				<span class='fas fa-bars'></span>
			</button>
			<div class='collapse navbar-collapse' id='main_navbar'>
				<ul class='navbar-nav'>
					<li class='nav-item dropdown '>
						<a class='nav-link' data-toggle='dropdown'  href='#' >
							<span class='fas fa-home' title="Home"></span>
	<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Home</span>
						</a>
						<ul class='dropdown-menu'>
							<li class='nav-item'><a class='nav-link' href='/core/users/user_edit.php?id=user' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Account Settings</a></li>
							<li class='nav-item'><a class='nav-link' href='/core/dashboard/' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Dashboard</a></li>
							<li class='nav-item'><a class='nav-link' href='/logout.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Logout</a></li>
						</ul>
					</li>
					<li class='nav-item dropdown '>
						<a class='nav-link' data-toggle='dropdown'  href='#' >
							<span class='fas fa-user' title="Accounts"></span>
	<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Accounts</span>
						</a>
						<ul class='dropdown-menu'>
							<li class='nav-item'><a class='nav-link' href='/app/devices/devices.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Devices</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/extensions/extensions.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Extensions</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/gateways/gateways.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Gateways</a></li>
							<li class='nav-item'><a class='nav-link' href='/core/users/users.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Users</a></li>
						</ul>
					</li>
					<li class='nav-item dropdown '>
						<a class='nav-link' data-toggle='dropdown'  href='#' >
							<span class='fas fa-exchange-alt' title="Dialplan"></span>
	<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Dialplan</span>
						</a>
						<ul class='dropdown-menu'>
							<li class='nav-item'><a class='nav-link' href='/app/destinations/destinations.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Destinations</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/dialplans/dialplans.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Dialplan Manager</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/dialplans/dialplans.php?app_uuid=c03b422e-13a8-bd1b-e42b-b6b9b4d27ce4' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Inbound Routes</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/dialplans/dialplans.php?app_uuid=8c914ec3-9fc0-8ab5-4cda-6c9288bdc9a3' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Outbound Routes</a></li>
						</ul>
					</li>
					<li class='nav-item dropdown '>
						<a class='nav-link' data-toggle='dropdown'  href='#' >
							<span class='fas fa-paper-plane' title="Applications"></span>
	<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Applications</span>
						</a>
						<ul class='dropdown-menu'>
							<li class='nav-item'><a class='nav-link' href='/app/bridges/bridges.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Bridges</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/call_block/call_block.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Block</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/call_broadcast/call_broadcast.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Broadcast</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/call_centers/call_center_queues.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Centers</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/xml_cdr/xml_cdr.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Detail Records</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/call_flows/call_flows.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Flows</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/call_forward/call_forward.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Forward</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/call_recordings/call_recordings.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Recordings</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/conference_centers/conference_centers.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Conference Centers</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/conference_controls/conference_controls.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Conference Controls</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/conference_profiles/conference_profiles.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Conference Profiles</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/conferences/conferences.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Conferences</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/contacts/contacts.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Contacts</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/fax/fax.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Fax Server</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/call_forward/call_forward.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Follow Me</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/gswave/index.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>GS Wave</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/ivr_menus/ivr_menus.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>IVR Menus</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/music_on_hold/music_on_hold.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Music on Hold</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/basic_operator_panel/index.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Operator Panel</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/phrases/phrases.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Phrases</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/dialplans/dialplans.php?app_uuid=16589224-c876-aeb3-f59f-523a1c0801f7' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Queues</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/recordings/recordings.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Recordings</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/ring_groups/ring_groups.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Ring Groups</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/streams/streams.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Streams</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/time_conditions/time_conditions.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Time Conditions</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/voicemails/voicemails.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Voicemail</a></li>
						</ul>
					</li>
					<li class='nav-item dropdown '>
						<a class='nav-link' data-toggle='dropdown'  href='#' >
							<span class='fas fa-chart-bar' title="Status"></span>
	<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Status</span>
						</a>
						<ul class='dropdown-menu'>
							<li class='nav-item'><a class='nav-link' href='/app/call_center_active/call_center_queue.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Active Call Center</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/calls_active/calls_active.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Active Calls</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/conferences_active/conferences_active.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Active Conferences</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/fifo_list/fifo_list.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Active Queues</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/call_centers/call_center_agent_status.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Agent Status</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/xml_cdr/xml_cdr_statistics.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>CDR Statistics</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/email_queue/email_queue.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Email Queue</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/event_guard/event_guard_logs.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Event Guard</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/xml_cdr/xml_cdr_extension_summary.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Extension Summary</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/fax_queue/fax_queue.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>FAX Queue</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/log_viewer/log_viewer.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Log Viewer</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/registrations/registrations.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Registrations</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/sip_status/sip_status.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>SIP Status</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/system/system.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>System Status</a></li>
							<li class='nav-item'><a class='nav-link' href='/core/user_logs/user_logs.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>User Logs</a></li>
						</ul>
					</li>
					<li class='nav-item dropdown '>
						<a class='nav-link' data-toggle='dropdown'  href='#' >
							<span class='fas fa-cog' title="Advanced"></span>
	<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Advanced</span>
						</a>
						<ul class='dropdown-menu'>
							<li class='nav-item'><a class='nav-link' href='/app/access_controls/access_controls.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Access Controls</a></li>
							<li class='nav-item'><a class='nav-link' href='/core/databases/databases.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Databases</a></li>
							<li class='nav-item'><a class='nav-link' href='/core/default_settings/default_settings.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Default Settings</a></li>
							<li class='nav-item'><a class='nav-link' href='/core/domains/domains.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Domains</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/email_templates/email_templates.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Email Templates</a></li>
							<li class='nav-item'><a class='nav-link' href='/core/groups/groups.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Group Manager</a></li>
							<li class='nav-item'><a class='nav-link' href='/core/menu/menu.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Menu Manager</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/modules/modules.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Modules</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/number_translations/number_translations.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Number Translations</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/sip_profiles/sip_profiles.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>SIP Profiles</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/settings/setting_edit.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Settings</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/database_transactions/database_transactions.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Transactions</a></li>
							<li class='nav-item'><a class='nav-link' href='/core/upgrade/index.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Upgrade</a></li>
							<li class='nav-item'><a class='nav-link' href='/app/vars/vars.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Variables</a></li>
						</ul>
					</li>
				</ul>
				<ul class='navbar-nav ml-auto'>
			<li class='nav-item'>
				<a class='header_domain' href='#' id='header_domain_selector_domain' title='Domain Selector [ESC]'><i class='fas fa-globe-americas fa-lg fa-fw' style='margin-top: 6px; margin-right: 5px;'></i>voip01&period;dial&period;opssbank&period;com&period;br</a>		</li>
				</ul>
			</div>
		</div>
	</nav>
	
					<div class='container-fluid' style='padding: 0;' align='center'>
	
							<div id='main_content'>
					<script>
	var Objs;
	
	function changeToInput(obj){
		tb=document.createElement('INPUT');
		tb.type='text';
		tb.name=obj.name;
		tb.setAttribute('class', 'formfld');
		tb.value=obj.options[obj.selectedIndex].value;
		tbb=document.createElement('INPUT');
		tbb.setAttribute('class', 'btn');
		tbb.setAttribute('style', 'margin-left: 4px;');
		tbb.type='button';
		tbb.value=$('<div />').html('&#9665;').text();
		tbb.objs=[obj,tb,tbb];
		tbb.onclick=function(){ Replace(this.objs); }
		obj.parentNode.insertBefore(tb,obj);
		obj.parentNode.insertBefore(tbb,obj);
		obj.parentNode.removeChild(obj);
	}
	
	function Replace(obj){
		obj[2].parentNode.insertBefore(obj[0],obj[2]);
		obj[0].parentNode.removeChild(obj[1]);
		obj[0].parentNode.removeChild(obj[2]);
	}
	</script>
	
	<form name='frm' id='frm' method='post'>
	<div class='action_bar' id='action_bar'>
		<div class='heading'><b>Call Center Queue</b></div>
		<div class='actions'>
	<a href='call_center_queues.php' target='_self' style='margin-right: 15px; ' ><button type='button' id='btn_back' alt='Back' title='Back' class='btn btn-default ' ><span class='fas fa-step-backward fa-fw'></span><span class='button-label  pad'>Back</span></button></a><a href='cmd.php?cmd=reload&id=988348b3-465d-449c-8137-ec3031941c84' target='_self' ><button type='button' alt='Reload' title='Reload' class='btn btn-default ' ><span class='fas fa-redo-alt fa-fw'></span><span class='button-label  pad'>Reload</span></button></a><a href='/app/call_center_active/call_center_active.php?queue_name=988348b3-465d-449c-8137-ec3031941c84' target='_self' style='margin-right: 15px; ' ><button type='button' alt='View' title='View' class='btn btn-default ' ><span class='fas fa-eye fa-fw'></span><span class='button-label  pad'>View</span></button></a><button type='submit' id='btn_save' alt='Save' title='Save' class='btn btn-default ' ><span class='fas fa-bolt fa-fw'></span><span class='button-label  pad'>Save</span></button>	</div>
		<div style='clear: both;'></div>
	</div>
	<table width='100%' border='0' cellpadding='0' cellspacing='0'>
	<tr>
	<td width='30%' class='vncellreq' valign='top' align='left' nowrap>
		Queue Name
	</td>
	<td width='70%' class='vtable' align='left'>
		<input class='formfld' type='text' name='queue_name' maxlength='255' value="Teste p&eacute;ricles" required='required'>
	<br />
	Enter the queue name.
	</td>
	</tr>
	<tr>
	<td class='vncellreq' valign='top' align='left' nowrap>
		Extension
	</td>
	<td class='vtable' align='left'>
		<input class='formfld' type='number' name='queue_extension' maxlength='255' min='0' step='1' value="200115" required='required'>
	<br />
	Enter the extension number.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap='nowrap'>
		Greeting
	</td>
	<td class='vtable' align='left'>
	<select name='queue_greeting' class='formfld' style='width: 200px;' onchange='changeToInput(this);'>
		<option value=''></option>
	<optgroup label=Miscellaneous>
		<option value='say&colon;'>say</option>
		<option value='tone&lowbar;stream&colon;'>tone&lowbar;stream</option>
	</optgroup>
		</select>
	<br />
	Select the desired Greeting.
	</td>
	</tr>
	<tr>
	<td class='vncellreq' valign='top' align='left' nowrap>
		Strategy
	</td>
	<td class='vtable' align='left'>
		<select class='formfld' name='queue_strategy'>
		<option value='ring-all'>Ring All</option>
		<option value='longest-idle-agent'>Longest Idle Agent</option>
		<option value='round-robin' selected='selected'>Round Robin</option>
		<option value='top-down'>Top Down</option>
		<option value='agent-with-least-talk-time'>Agent With Least Talk Time</option>
		<option value='agent-with-fewest-calls'>Agent With Fewest Calls</option>
		<option value='sequentially-by-agent-order'>Sequentially By Agent Order</option>
		<option value='sequentially-by-next-agent-order'>Sequentially By Next Agent Order</option>
		<option value='random'>Random</option>
		</select>
	<br />
	Select the queue ring strategy.
	</td>
	</tr>
	<tr>	<td class='vncell' valign='top'>Agents</td>	<td class='vtable' align='left'>			<table border='0' cellpadding='0' cellspacing='0'>
				<tr>
					<td class='vtable'>Agent Name</td>
					<td class='vtable' style='text-align: center;'>Tier Level</td>
					<td class='vtable' style='text-align: center;'>Tier Position</td>
					<td></td>
				</tr>
		<tr>
			<td class=''>				<input name='call_center_tiers[0][call_center_tier_uuid]' type='hidden' value="5fbd5a8b-c315-4c41-9713-3ed51e22c403">
					<select name="call_center_tiers[0][call_center_agent_uuid]" class="formfld" style="width: 200px">
					<option value="7caa54c4-9899-46fb-8177-ebc1f83513d2">1638</option>
					</select>		</td>
			<td class='' style='text-align: center;'>				 <select name="call_center_tiers[0][tier_level]" class="formfld">
					<option value="0" selected>0</option>
					<option value="1" >1</option>
					<option value="2" >2</option>
					<option value="3" >3</option>
					<option value="4" >4</option>
					<option value="5" >5</option>
					<option value="6" >6</option>
					<option value="7" >7</option>
					<option value="8" >8</option>
					<option value="9" >9</option>
					</select>
			</td>
			<td class='' style='text-align: center;'>
					<select name="call_center_tiers[0][tier_position]" class="formfld">
					<option value="0" selected>0</option>
					<option value="1" >1</option>
					<option value="2" >2</option>
					<option value="3" >3</option>
					<option value="4" >4</option>
					<option value="5" >5</option>
					<option value="6" >6</option>
					<option value="7" >7</option>
					<option value="8" >8</option>
					<option value="9" >9</option>
					</select>
			</td>
			<td class=''>			<a href="call_center_queue_edit.php?id=988348b3-465d-449c-8137-ec3031941c84&call_center_tier_uuid=5fbd5a8b-c315-4c41-9713-3ed51e22c403&a=delete" alt="Delete" onclick="return confirm('Do you really want to DELETE this?');"><button type='button' class='btn btn-default list_control_icon'><span class='fas fa-minus'></span></button></a>		</td>
		</tr>
		<tr>
			<td class=''>				<input name='call_center_tiers[1][call_center_tier_uuid]' type='hidden' value="3885cff9-93c0-42ae-b6ca-1f56b5e77385">
					<select name="call_center_tiers[1][call_center_agent_uuid]" class="formfld" style="width: 200px">
						<option value=""></option>
					<option value="fd7b879d-a2dd-4ef7-9b05-06156a4f80b7">1601</option>
					<option value="6d66dadc-2d38-46e3-ba15-aa113ffa4892">1603</option>
					<option value="c2c0722e-0550-4caa-852e-b765791ef931">1604</option>
					<option value="6a256f0c-84a4-494a-8f66-2913e59779b6">1606</option>
					<option value="a342da12-0d53-4362-8238-417c7a50368c">1607</option>
					<option value="9557f6cf-785d-4e6e-ada6-b94910190556">1609</option>
					<option value="ef50191d-d238-4e44-a29c-db1c22634a21">1610</option>
					<option value="0e96e013-8199-47fd-8268-6c1dcd20aa96">1611</option>
					<option value="d1499462-3461-4d93-9441-72b3c24f857c">1612</option>
					<option value="78352cbe-c627-41cf-a212-a3fc8b8c6260">1613</option>
					<option value="305dbdf8-4546-4053-aa3e-fc268d5f9c91">1614</option>
					<option value="825a6bf5-809f-49fb-a197-7caffa87fd2d">1615</option>
					<option value="98d94ddc-8496-4b95-ae91-9ceee7380083">1616</option>
					<option value="b623066e-8587-4742-8727-a2228c3375d9">1617</option>
					<option value="b6a7e90d-2e94-4967-be9b-48d9c89db12e">1618</option>
					<option value="082eb3b0-fae5-4257-b538-be5e1f9156ca">1619</option>
					<option value="6d2ea436-325d-492c-9143-2b7298ad190a">1626</option>
					<option value="dfabcc60-538e-49dc-8deb-9ab5dedd8c36">1627</option>
					<option value="24d82d98-e94a-4c7f-b64a-feb05261ff6e">1628</option>
					<option value="a6862035-0524-48ee-921c-9b30dc31c2cf">1629</option>
					<option value="3961b4d6-210e-45e9-ac4f-a2ecb66913a5">1630</option>
					<option value="353d6cb2-dbec-44c1-995a-fe5d86e6e277">1631</option>
					<option value="30c7a03d-c175-431d-8c18-d74e1975cd0d">1632</option>
					<option value="fc9a2911-2a09-48ff-b89b-bda47b21998a">1634</option>
					<option value="6fe18ce4-8327-46a8-baab-0f228bcf5649">1635</option>
					<option value="6214f0c8-395c-4c59-87b5-002702195c51">1636</option>
					<option value="f8fba913-a39f-4ed2-94e5-6a51b931f80a">1637</option>
					<option value="7caa54c4-9899-46fb-8177-ebc1f83513d2">1638</option>
					<option value="2e92ca8c-9c72-40d3-9444-04da4bf50a6a">1639</option>
					<option value="f998c21b-f324-47c9-9624-a47876cf0a91">1641</option>
					<option value="d3aa8b8e-54e0-42e9-9f82-65e4fa996796">1642</option>
					<option value="e1f442e1-cbd9-4da9-a9db-2beaf61898b1">1644</option>
					<option value="52b282f8-6499-4b53-a92d-7407c2a601d0">1645</option>
					<option value="50f623e0-3092-428d-bc50-03f086867a82">1646</option>
					<option value="11940fa2-4bd9-433e-b4ad-5b8fb264422d">1647</option>
					<option value="d144d59d-da87-4640-8df5-ab44a4dee372">1649</option>
					<option value="13d3de44-422b-4cf7-aaca-dbb800dfd762">1650</option>
					<option value="864bec58-5173-40b4-8799-5995bb19cb2c">1651</option>
					<option value="6c504577-0c53-4aa5-8da1-c5ca00ebdcd6">1652</option>
					<option value="ac105f6c-6655-46b1-8adf-67e75b3fc6d5">1653</option>
					<option value="264d0e22-a4c8-42e9-947f-a67fab9f2fb9">1654</option>
					<option value="21ee25df-ad7e-4e82-83e7-379670220cdb">1655</option>
					<option value="5baa539e-ab37-4633-9351-42eff7c15edc">1656</option>
					<option value="91f77326-847d-4ed2-8bd1-92ee71180fd3">1657</option>
					<option value="2cdfd4fb-db22-4b31-93b0-febe5dfc61e3">1658</option>
					<option value="598ebf17-4471-475c-8451-7b25d3e1383e">1659</option>
					<option value="dabf7ba9-6d74-48df-b533-dda7219cea07">1660</option>
					<option value="05cf3919-cec9-4b12-9525-b492fa426da3">1661</option>
					<option value="92569c7d-f142-440f-a885-a86c72ad0936">1662</option>
					<option value="015cdc3e-26e7-4347-a2d7-483f0b88ec01">1663</option>
					<option value="1d20eec5-4f2f-43c8-bf2c-d06fb2e272f8">1664</option>
					<option value="97227560-247d-49b7-99c5-c19e3cac9552">1665</option>
					<option value="a958700d-5c31-48d7-9f87-e05cbebfc6c5">1666</option>
					<option value="1c063326-64e2-457c-876d-f536a8ba23a2">1668</option>
					<option value="19deda13-ef7f-43b7-a8c6-3d3f9601ca3c">1669</option>
					<option value="99b094b6-5e08-4472-adc9-853c604be499">1670</option>
					<option value="1265e427-d385-4a3e-a6dc-0093880cd1e5">1671</option>
					<option value="d796a830-cf23-4589-9b07-e610c6d2e7d8">1672</option>
					<option value="dbe0646d-7afb-4ba3-91b0-774952d1021b">1673</option>
					<option value="779905d7-fa83-412f-ad3e-433b53719835">1674</option>
					<option value="b3204faf-86f8-4530-97bd-a6aa22e4886c">1675</option>
					<option value="1bd17098-016f-44ab-b0d0-c6d42fbb944c">1676</option>
					<option value="f81bfa23-068e-41c1-b0b0-643652007a96">1677</option>
					<option value="898d2219-2288-40c6-b2bb-0ecf156cea95">1678</option>
					<option value="5b3d49a1-417d-46b6-b2f2-e6a2f97de078">1679</option>
					<option value="faca8f9b-4f10-403d-9e56-6ab0dae2acbd">1680</option>
					<option value="e9c17f2f-54e8-4a90-852e-fa7d5b1c4f1b">1681</option>
					<option value="9becac36-310b-4852-b464-062e88968c54">1682</option>
					<option value="ce2a697d-85e8-4bd7-a8b5-2f5b7906b883">1683</option>
					<option value="1dde0bf9-966a-40f4-8f02-68928048f0af">1684</option>
					<option value="77dadef4-03b9-4c8a-973e-2b912789a972">1685</option>
					<option value="d93b5abb-5dcf-4e40-af16-7338f5f64c3a">1686</option>
					<option value="92374956-dd15-4edb-a340-70139e0e941c">1687</option>
					<option value="e3923b66-6cf3-4da6-b1aa-caba339127c0">1688</option>
					<option value="567ae4fc-db87-4a4b-b405-265c50d55604">1689</option>
					<option value="7bd3f9e8-b321-4e92-a862-09d416ae1594">1690</option>
					<option value="38ad6832-537f-4f7b-af37-5b8f0f755359">1691</option>
					<option value="ac68ed3a-cf71-4d66-9577-3f6512fed3d1">1692</option>
					<option value="b7ad4660-4bd7-42a0-9223-f96bceac18b6">1693</option>
					<option value="9d3f1651-e68e-44b8-b425-2f697a0109af">1694</option>
					<option value="b70815c9-1672-414a-ae4c-25b39f54beff">1696</option>
					<option value="77cdd4d6-c15b-45c8-ad6c-671d8f9426b6">9998</option>
					<option value="b4ca43fc-d988-46d3-9cc6-2a8ecc5a0893">9999</option>
					</select>		</td>
			<td class='' style='text-align: center;'>				 <select name="call_center_tiers[1][tier_level]" class="formfld">
					<option value="0" >0</option>
					<option value="1" >1</option>
					<option value="2" >2</option>
					<option value="3" >3</option>
					<option value="4" >4</option>
					<option value="5" >5</option>
					<option value="6" >6</option>
					<option value="7" >7</option>
					<option value="8" >8</option>
					<option value="9" >9</option>
					</select>
			</td>
			<td class='' style='text-align: center;'>
					<select name="call_center_tiers[1][tier_position]" class="formfld">
					<option value="0" >0</option>
					<option value="1" >1</option>
					<option value="2" >2</option>
					<option value="3" >3</option>
					<option value="4" >4</option>
					<option value="5" >5</option>
					<option value="6" >6</option>
					<option value="7" >7</option>
					<option value="8" >8</option>
					<option value="9" >9</option>
					</select>
			</td>
			<td class=''>			<a href="call_center_queue_edit.php?id=988348b3-465d-449c-8137-ec3031941c84&call_center_tier_uuid=3885cff9-93c0-42ae-b6ca-1f56b5e77385&a=delete" alt="Delete" onclick="return confirm('Do you really want to DELETE this?');"><button type='button' class='btn btn-default list_control_icon'><span class='fas fa-minus'></span></button></a>		</td>
		</tr>
			</table>
			<br>
			Tiers assign agents to queues.
			<br />
		</td></tr><tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Music on Hold
	</td>
	<td class='vtable' align='left'>
	<select class='formfld' name='queue_moh_sound' id='queue_moh_sound' style='width: auto;'>
			<option value=''></option>
		<optgroup label='Music on Hold'>
			<option value='local_stream://default' >default</option>
		</optgroup>
		<optgroup label='Ringtones'>		<option value='${au-ring}'>au-ring</option>
			<option value='${be-ring}'>be-ring</option>
			<option value='${bong-ring}'>bong-ring</option>
			<option value='${ca-ring}'>ca-ring</option>
			<option value='${cn-ring}'>cn-ring</option>
			<option value='${cy-ring}'>cy-ring</option>
			<option value='${cz-ring}'>cz-ring</option>
			<option value='${de-ring}'>de-ring</option>
			<option value='${dk-ring}'>dk-ring</option>
			<option value='${dz-ring}'>dz-ring</option>
			<option value='${eg-ring}'>eg-ring</option>
			<option value='${fi-ring}'>fi-ring</option>
			<option value='${fr-ring}'>fr-ring</option>
			<option value='${hk-ring}'>hk-ring</option>
			<option value='${hu-ring}'>hu-ring</option>
			<option value='${il-ring}'>il-ring</option>
			<option value='${in-ring}'>in-ring</option>
			<option value='${it-ring}'>it-ring</option>
			<option value='${jp-ring}'>jp-ring</option>
			<option value='${ko-ring}'>ko-ring</option>
			<option value='${pk-ring}'>pk-ring</option>
			<option value='${pl-ring}'>pl-ring</option>
			<option value='${pt-ring}' selected="selected">pt-ring</option>
			<option value='${ro-ring}'>ro-ring</option>
			<option value='${rs-ring}'>rs-ring</option>
			<option value='${ru-ring}'>ru-ring</option>
			<option value='${sa-ring}'>sa-ring</option>
			<option value='${tr-ring}'>tr-ring</option>
			<option value='${uk-ring}'>uk-ring</option>
			<option value='${us-ring}'>us-ring</option>
			<option value='silence'>Silence</option>
		</optgroup>
		<optgroup label='Tones'>		<option value='${bong-us-tone}'>bong-us-tone</option>
			<option value='${busy-au-tone}'>busy-au-tone</option>
			<option value='${busy-us-tone}'>busy-us-tone</option>
			<option value='${vacant-uk-tone}'>vacant-uk-tone</option>
			<option value='${vacant-us-tone}'>vacant-us-tone</option>
		</optgroup>
	</select>
	<br />
	Select the desired hold music.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Record
	</td>
	<td class='vtable' align='left'>
		<select class='formfld' name='queue_record_template'>
		<option value='&sol;var&sol;lib&sol;freeswitch&sol;recordings&sol;voip01&period;dial&period;opssbank&period;com&period;br&sol;archive&sol;&dollar;&lbrace;strftime&lpar;&percnt;Y&rpar;&rcub;&sol;&dollar;&lbrace;strftime&lpar;&percnt;b&rpar;&rcub;&sol;&dollar;&lbrace;strftime&lpar;&percnt;d&rpar;&rcub;&sol;&dollar;&lbrace;uuid&rcub;&period;&dollar;&lbrace;record&lowbar;ext&rcub;' selected='selected' >True</option>
		<option value=''>False</option>
		</select>
	<br />
	Save the recording.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Time Base Score
	</td>
	<td class='vtable' align='left'>
		<select class='formfld' name='queue_time_base_score'>
		<option value='system' selected='selected' >System</option>
		<option value='queue'>Queue</option>
		</select>
	<br />
	Select the time base score.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Time Base Score Seconds
	</td>
	<td class='vtable' align='left'>
	  <input class='formfld' type='number' name='queue_time_base_score_sec' maxlength='255' min='0' step='1' value=''>
	<br />
	Set the time base score in seconds. Higher numbers mean higher priority.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Max Wait Time
	</td>
	<td class='vtable' align='left'>
	  <input class='formfld' type='number' name='queue_max_wait_time' maxlength='255' min='0' step='1' value='0'>
	<br />
	Enter the max wait time.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Max Wait Time with No Agent
	</td>
	<td class='vtable' align='left'>
	  <input class='formfld' type='number' name='queue_max_wait_time_with_no_agent' maxlength='255' min='0' step='1' value='90'>
	<br />
	Enter the max wait time with no agent.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Max Wait Time with No Agent Time Reached
	</td>
	<td class='vtable' align='left'>
	  <input class='formfld' type='number' name='queue_max_wait_time_with_no_agent_time_reached' maxlength='255' min='0' step='1' value='30'>
	<br />
	Enter the max wait time with no agent time reached.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Timeout Action
	</td>
	<td class='vtable' align='left'>
	
	<script>
	var Objs;
	
	function changeToInputqueue_timeout_action(obj){
		tb=document.createElement('INPUT');
		tb.type='text';
		tb.name=obj.name;
		tb.className='formfld';
		tb.setAttribute('id', 'queue_timeout_action');
		tb.setAttribute('style', 'width: 200px;');
		tb.value=obj.options[obj.selectedIndex].value;
		document.getElementById('btn_select_to_input_queue_timeout_action').style.visibility = 'hidden';
		tbb=document.createElement('INPUT');
		tbb.setAttribute('class', 'btn');
		tbb.setAttribute('style', 'margin-left: 4px;');
		tbb.type='button';
		tbb.value=$('<div />').html('&#9665;').text();
		tbb.objs=[obj,tb,tbb];
		tbb.onclick=function(){ Replacequeue_timeout_action(this.objs); }
		obj.parentNode.insertBefore(tb,obj);
		obj.parentNode.insertBefore(tbb,obj);
		obj.parentNode.removeChild(obj);
		Replacequeue_timeout_action(this.objs);
	}
	
	function Replacequeue_timeout_action(obj){
		obj[2].parentNode.insertBefore(obj[0],obj[2]);
		obj[0].parentNode.removeChild(obj[1]);
		obj[0].parentNode.removeChild(obj[2]);
		document.getElementById('btn_select_to_input_queue_timeout_action').style.visibility = 'visible';
	}
	</script>
	
		<select name='queue_timeout_action' id='queue_timeout_action' class='formfld' style='width: 200px;' onchange="">
				<option value=''></option>
			<optgroup label='Call Center'>
				<option value='transfer&colon;55550 XML voip01&period;dial&period;opssbank&period;com&period;br' >55550 ATENDIMENTO</option>
				<option value='transfer&colon;808080 XML voip01&period;dial&period;opssbank&period;com&period;br' >808080 AT&lowbar;CRM</option>
				<option value='transfer&colon;200110 XML voip01&period;dial&period;opssbank&period;com&period;br' >200110 Admin Opss</option>
				<option value='transfer&colon;200111 XML voip01&period;dial&period;opssbank&period;com&period;br' >200111 CAMPANHAS DE MKT</option>
				<option value='transfer&colon;66666 XML voip01&period;dial&period;opssbank&period;com&period;br' >66666 CRM</option>
				<option value='transfer&colon;200100 XML voip01&period;dial&period;opssbank&period;com&period;br' >200100 Cavalcante &amp; Mendes</option>
				<option value='transfer&colon;200101 XML voip01&period;dial&period;opssbank&period;com&period;br' >200101 Cedro 7 Seguros</option>
				<option value='transfer&colon;200102 XML voip01&period;dial&period;opssbank&period;com&period;br' >200102 Davi Miret</option>
				<option value='transfer&colon;200104 XML voip01&period;dial&period;opssbank&period;com&period;br' >200104 FIRMASEG</option>
				<option value='transfer&colon;200105 XML voip01&period;dial&period;opssbank&period;com&period;br' >200105 FP 2</option>
				<option value='transfer&colon;200103 XML voip01&period;dial&period;opssbank&period;com&period;br' >200103 Felice Adm</option>
				<option value='transfer&colon;200106 XML voip01&period;dial&period;opssbank&period;com&period;br' >200106 GHGM 1</option>
				<option value='transfer&colon;200112 XML voip01&period;dial&period;opssbank&period;com&period;br' >200112 Home Office</option>
				<option value='transfer&colon;200107 XML voip01&period;dial&period;opssbank&period;com&period;br' >200107 Inonvax</option>
				<option value='transfer&colon;200108 XML voip01&period;dial&period;opssbank&period;com&period;br' >200108 Ivaro Matos</option>
				<option value='transfer&colon;200109 XML voip01&period;dial&period;opssbank&period;com&period;br' >200109 Nascional 1</option>
				<option value='transfer&colon;200113 XML voip01&period;dial&period;opssbank&period;com&period;br' >200113 Opss Lab</option>
				<option value='transfer&colon;200114 XML voip01&period;dial&period;opssbank&period;com&period;br' >200114 Opss Team</option>
				<option value='transfer&colon;200116 XML voip01&period;dial&period;opssbank&period;com&period;br' >200116 Prevseguros</option>
				<option value='transfer&colon;200118 XML voip01&period;dial&period;opssbank&period;com&period;br' >200118 RF Feijo</option>
				<option value='transfer&colon;200117 XML voip01&period;dial&period;opssbank&period;com&period;br' >200117 RIB COR</option>
				<option value='transfer&colon;200119 XML voip01&period;dial&period;opssbank&period;com&period;br' >200119 Ronaldo Alex Seguros</option>
				<option value='transfer&colon;200115 XML voip01&period;dial&period;opssbank&period;com&period;br' >200115 Teste p&eacute;ricles</option>
				<option value='transfer&colon;200220 XML voip01&period;dial&period;opssbank&period;com&period;br' >200220 Victor Hugo Ferreira</option>
				<option value='transfer&colon;200221 XML voip01&period;dial&period;opssbank&period;com&period;br' >200221 Wilson Cor</option>
				<option value='transfer&colon;200170 XML voip01&period;dial&period;opssbank&period;com&period;br' >200170 c&lowbar;170</option>
			</optgroup>
			<optgroup label='Call Groups'>
				<option value='bridge&colon;group&sol;sales&commat;voip01&period;dial&period;opssbank&period;com&period;br' >sales</option>
				<option value='bridge&colon;group&sol;support&commat;voip01&period;dial&period;opssbank&period;com&period;br' >support</option>
			</optgroup>
			<optgroup label='Extensions'>
				<option value='transfer&colon;1467 XML voip01&period;dial&period;opssbank&period;com&period;br' >1467 Rogerio Pessoa</option>
				<option value='transfer&colon;1601 XML voip01&period;dial&period;opssbank&period;com&period;br' >1601 Danilo Machado</option>
				<option value='transfer&colon;1602 XML voip01&period;dial&period;opssbank&period;com&period;br' >1602 1602WEBRTC</option>
				<option value='transfer&colon;1603 XML voip01&period;dial&period;opssbank&period;com&period;br' >1603 1603descricao</option>
				<option value='transfer&colon;1604 XML voip01&period;dial&period;opssbank&period;com&period;br' >1604 Marcelo Palmieri</option>
				<option value='transfer&colon;1606 XML voip01&period;dial&period;opssbank&period;com&period;br' >1606 Gustavo Felipe</option>
				<option value='transfer&colon;1607 XML voip01&period;dial&period;opssbank&period;com&period;br' >1607 Franciane Kelly</option>
				<option value='transfer&colon;1608 XML voip01&period;dial&period;opssbank&period;com&period;br' >1608 Leonardo Mendes</option>
				<option value='transfer&colon;1609 XML voip01&period;dial&period;opssbank&period;com&period;br' >1609 Wesley Jose</option>
				<option value='transfer&colon;1610 XML voip01&period;dial&period;opssbank&period;com&period;br' >1610 Marcos Machado</option>
				<option value='transfer&colon;1611 XML voip01&period;dial&period;opssbank&period;com&period;br' >1611 Vitor Junior</option>
				<option value='transfer&colon;1612 XML voip01&period;dial&period;opssbank&period;com&period;br' >1612 Alexandre Dias</option>
				<option value='transfer&colon;1613 XML teste&lowbar;algar' >1613 1613</option>
				<option value='transfer&colon;1614 XML voip01&period;dial&period;opssbank&period;com&period;br' >1614 Gabriela Araujo</option>
				<option value='transfer&colon;1615 XML voip01&period;dial&period;opssbank&period;com&period;br' >1615 Michell Augusto</option>
				<option value='transfer&colon;1616 XML voip01&period;dial&period;opssbank&period;com&period;br' >1616 Mateus Assis</option>
				<option value='transfer&colon;1617 XML voip01&period;dial&period;opssbank&period;com&period;br' >1617 1617</option>
				<option value='transfer&colon;1618 XML voip01&period;dial&period;opssbank&period;com&period;br' >1618 Andre Ventura</option>
				<option value='transfer&colon;1619 XML voip01&period;dial&period;opssbank&period;com&period;br' >1619 1619</option>
				<option value='transfer&colon;1620 XML voip01&period;dial&period;opssbank&period;com&period;br' >1620 POSITIVA ESTHER MARTINS</option>
				<option value='transfer&colon;1621 XML voip01&period;dial&period;opssbank&period;com&period;br' >1621 POSITIVA GABRIELA SANTOS</option>
				<option value='transfer&colon;1622 XML voip01&period;dial&period;opssbank&period;com&period;br' >1622 POSITIVA JULIA BEATRIZ</option>
				<option value='transfer&colon;1623 XML voip01&period;dial&period;opssbank&period;com&period;br' >1623 POSITIVA MARIA GABRIELLI</option>
				<option value='transfer&colon;1624 XML voip01&period;dial&period;opssbank&period;com&period;br' >1624 POSITIVA IEDA CRISTINA</option>
				<option value='transfer&colon;1625 XML voip01&period;dial&period;opssbank&period;com&period;br' >1625 Positiva Victoria</option>
				<option value='transfer&colon;1626 XML voip01&period;dial&period;opssbank&period;com&period;br' >1626 Biane Tatielly</option>
				<option value='transfer&colon;1627 XML voip01&period;dial&period;opssbank&period;com&period;br' >1627 Marcos Machado</option>
				<option value='transfer&colon;1628 XML voip01&period;dial&period;opssbank&period;com&period;br' >1628 Regiane Oliveira</option>
				<option value='transfer&colon;1629 XML voip01&period;dial&period;opssbank&period;com&period;br' >1629 Stephanie Ramos</option>
				<option value='transfer&colon;1630 XML voip01&period;dial&period;opssbank&period;com&period;br' >1630 Vitor Junio</option>
				<option value='transfer&colon;1631 XML voip01&period;dial&period;opssbank&period;com&period;br' >1631 Viviane Adame</option>
				<option value='transfer&colon;1632 XML voip01&period;dial&period;opssbank&period;com&period;br' >1632 Diogo Machado</option>
				<option value='transfer&colon;1633 XML voip01&period;dial&period;opssbank&period;com&period;br' >1633 Rodrigo Mattos</option>
				<option value='transfer&colon;1634 XML voip01&period;dial&period;opssbank&period;com&period;br' >1634 Augusto Nascimento</option>
				<option value='transfer&colon;1635 XML voip01&period;dial&period;opssbank&period;com&period;br' >1635 Douglas Machado</option>
				<option value='transfer&colon;1636 XML voip01&period;dial&period;opssbank&period;com&period;br' >1636 Eduarda Guedes</option>
				<option value='transfer&colon;1637 XML voip01&period;dial&period;opssbank&period;com&period;br' >1637 1637</option>
				<option value='transfer&colon;1641 XML voip01&period;dial&period;opssbank&period;com&period;br' >1641 Debora Menezes</option>
				<option value='transfer&colon;1642 XML voip01&period;dial&period;opssbank&period;com&period;br' >1642 Inovax David Jesus</option>
				<option value='transfer&colon;1643 XML voip01&period;dial&period;opssbank&period;com&period;br' >1643 Inovax Julia de Souza</option>
				<option value='transfer&colon;1644 XML voip01&period;dial&period;opssbank&period;com&period;br' >1644 Ana Julia</option>
				<option value='transfer&colon;1645 XML voip01&period;dial&period;opssbank&period;com&period;br' >1645 Lorraine Barbosa</option>
				<option value='transfer&colon;1646 XML voip01&period;dial&period;opssbank&period;com&period;br' >1646 Natalia Duarte</option>
				<option value='transfer&colon;1647 XML voip01&period;dial&period;opssbank&period;com&period;br' >1647 Ronaldo Alex</option>
				<option value='transfer&colon;1649 XML voip01&period;dial&period;opssbank&period;com&period;br' >1649 1649</option>
				<option value='transfer&colon;1650 XML voip01&period;dial&period;opssbank&period;com&period;br' >1650 1650</option>
				<option value='transfer&colon;1651 XML voip01&period;dial&period;opssbank&period;com&period;br' >1651 1651</option>
				<option value='transfer&colon;1652 XML voip01&period;dial&period;opssbank&period;com&period;br' >1652 1652</option>
				<option value='transfer&colon;1653 XML voip01&period;dial&period;opssbank&period;com&period;br' >1653 1653</option>
				<option value='transfer&colon;1654 XML voip01&period;dial&period;opssbank&period;com&period;br' >1654 1654</option>
				<option value='transfer&colon;1655 XML voip01&period;dial&period;opssbank&period;com&period;br' >1655 1655</option>
				<option value='transfer&colon;1656 XML voip01&period;dial&period;opssbank&period;com&period;br' >1656 Gleiciane Rodrigues</option>
				<option value='transfer&colon;1657 XML voip01&period;dial&period;opssbank&period;com&period;br' >1657 Vivian Fernandes</option>
				<option value='transfer&colon;1658 XML voip01&period;dial&period;opssbank&period;com&period;br' >1658 Erilaine&nbsp;Santos</option>
				<option value='transfer&colon;1659 XML voip01&period;dial&period;opssbank&period;com&period;br' >1659 Fladielle Santos</option>
				<option value='transfer&colon;1660 XML voip01&period;dial&period;opssbank&period;com&period;br' >1660 Ana Neres</option>
				<option value='transfer&colon;1661 XML voip01&period;dial&period;opssbank&period;com&period;br' >1661 1661</option>
				<option value='transfer&colon;1662 XML voip01&period;dial&period;opssbank&period;com&period;br' >1662 1662</option>
				<option value='transfer&colon;1663 XML voip01&period;dial&period;opssbank&period;com&period;br' >1663 1663</option>
				<option value='transfer&colon;1664 XML voip01&period;dial&period;opssbank&period;com&period;br' >1664 1664</option>
				<option value='transfer&colon;1665 XML voip01&period;dial&period;opssbank&period;com&period;br' >1665 1665</option>
				<option value='transfer&colon;1666 XML voip01&period;dial&period;opssbank&period;com&period;br' >1666 1666</option>
				<option value='transfer&colon;1667 XML voip01&period;dial&period;opssbank&period;com&period;br' >1667 Joao Victor</option>
				<option value='transfer&colon;1668 XML voip01&period;dial&period;opssbank&period;com&period;br' >1668 Maria Eduarda</option>
				<option value='transfer&colon;1669 XML voip01&period;dial&period;opssbank&period;com&period;br' >1669 Bruna Antonia</option>
				<option value='transfer&colon;1670 XML voip01&period;dial&period;opssbank&period;com&period;br' >1670 Daniel Kesley</option>
				<option value='transfer&colon;1671 XML voip01&period;dial&period;opssbank&period;com&period;br' >1671 Cecilia Aparecida</option>
				<option value='transfer&colon;1672 XML voip01&period;dial&period;opssbank&period;com&period;br' >1672 Elisangela Alves</option>
				<option value='transfer&colon;1673 XML voip01&period;dial&period;opssbank&period;com&period;br' >1673 Nadia Priscila</option>
				<option value='transfer&colon;1674 XML voip01&period;dial&period;opssbank&period;com&period;br' >1674 Analice Soares</option>
				<option value='transfer&colon;1676 XML voip01&period;dial&period;opssbank&period;com&period;br' >1676 1676</option>
				<option value='transfer&colon;1677 XML voip01&period;dial&period;opssbank&period;com&period;br' >1677 1677</option>
				<option value='transfer&colon;1678 XML voip01&period;dial&period;opssbank&period;com&period;br' >1678 1678</option>
				<option value='transfer&colon;1679 XML voip01&period;dial&period;opssbank&period;com&period;br' >1679 1679</option>
				<option value='transfer&colon;1680 XML voip01&period;dial&period;opssbank&period;com&period;br' >1680 1680</option>
				<option value='transfer&colon;1681 XML voip01&period;dial&period;opssbank&period;com&period;br' >1681 1681</option>
				<option value='transfer&colon;1682 XML voip01&period;dial&period;opssbank&period;com&period;br' >1682 1682</option>
				<option value='transfer&colon;1683 XML voip01&period;dial&period;opssbank&period;com&period;br' >1683 1683</option>
				<option value='transfer&colon;1684 XML voip01&period;dial&period;opssbank&period;com&period;br' >1684 1684</option>
				<option value='transfer&colon;1685 XML voip01&period;dial&period;opssbank&period;com&period;br' >1685 1685</option>
				<option value='transfer&colon;1686 XML voip01&period;dial&period;opssbank&period;com&period;br' >1686 1686</option>
				<option value='transfer&colon;1687 XML voip01&period;dial&period;opssbank&period;com&period;br' >1687 1687</option>
				<option value='transfer&colon;1688 XML voip01&period;dial&period;opssbank&period;com&period;br' >1688 1688</option>
				<option value='transfer&colon;1689 XML voip01&period;dial&period;opssbank&period;com&period;br' >1689 1689</option>
				<option value='transfer&colon;1690 XML voip01&period;dial&period;opssbank&period;com&period;br' >1690 1690</option>
				<option value='transfer&colon;1691 XML voip01&period;dial&period;opssbank&period;com&period;br' >1691 1691</option>
				<option value='transfer&colon;1692 XML voip01&period;dial&period;opssbank&period;com&period;br' >1692 1692</option>
				<option value='transfer&colon;1693 XML voip01&period;dial&period;opssbank&period;com&period;br' >1693 1693</option>
				<option value='transfer&colon;1694 XML voip01&period;dial&period;opssbank&period;com&period;br' >1694 1694</option>
				<option value='transfer&colon;1695 XML voip01&period;dial&period;opssbank&period;com&period;br' >1695 1695</option>
				<option value='transfer&colon;1696 XML voip01&period;dial&period;opssbank&period;com&period;br' >1696 1696</option>
				<option value='transfer&colon;1697 XML voip01&period;dial&period;opssbank&period;com&period;br' >1697</option>
				<option value='transfer&colon;7777 XML escritorio' >7777 CHARLESTON</option>
				<option value='transfer&colon;8887 XML escritorio' >8887 RAMAL CALL CENTER</option>
				<option value='transfer&colon;8888 XML escritorio' >8888 RAMAL ESCRITORIO</option>
				<option value='transfer&colon;9996 XML discador&lowbar;externo' >9996</option>
				<option value='transfer&colon;9997 XML voip01&period;dial&period;opssbank&period;com&period;br' >9997</option>
				<option value='transfer&colon;9998 XML voip01&period;dial&period;opssbank&period;com&period;br' >9998</option>
				<option value='transfer&colon;9999 XML voip01&period;dial&period;opssbank&period;com&period;br' >9999</option>
			</optgroup>
			<optgroup label='Tones'>
				<option value='playback&colon;tone&lowbar;stream&colon;&sol;&sol;v&equals;-7&semi;&percnt;&lpar;100&comma;0&comma;941&period;0&comma;1477&period;0&rpar;&semi;v&equals;-7&semi;&gt;&equals;2&semi;&plus;&equals;&period;1&semi;&percnt;&lpar;1400&comma;0&comma;350&comma;440&rpar;' >bong-us-tone</option>
				<option value='playback&colon;tone&lowbar;stream&colon;&sol;&sol;v&equals;-13&semi;&percnt;&lpar;375&comma;375&comma;420&rpar;&semi;v&equals;-23&semi;&percnt;&lpar;375&comma;375&comma;420&rpar;' >busy-au-tone</option>
				<option value='playback&colon;tone&lowbar;stream&colon;&sol;&sol;&percnt;&lpar;500&comma;500&comma;480&comma;620&rpar;' >busy-us-tone</option>
				<option value='playback&colon;tone&lowbar;stream&colon;&sol;&sol;&percnt;&lpar;330&comma;15&comma;950&rpar;&semi;&percnt;&lpar;330&comma;15&comma;1400&rpar;&semi;&percnt;&lpar;330&comma;1000&comma;1800&rpar;' >vacant-uk-tone</option>
				<option value='playback&colon;tone&lowbar;stream&colon;&sol;&sol;&percnt;&lpar;274&comma;0&comma;913&period;8&rpar;&semi;&percnt;&lpar;274&comma;0&comma;1370&period;6&rpar;&semi;&percnt;&lpar;380&comma;0&comma;1776&period;7&rpar;' >vacant-us-tone</option>
			</optgroup>
			<optgroup label='Voicemails'>
				<option value='transfer&colon;&ast;991608 XML voip01&period;dial&period;opssbank&period;com&period;br' >1608 Leonardo Mendes &#9993</option>
				<option value='transfer&colon;&ast;991613 XML voip01&period;dial&period;opssbank&period;com&period;br' >1613 1613 &#9993</option>
				<option value='transfer&colon;&ast;991638 XML voip01&period;dial&period;opssbank&period;com&period;br' >1638 Camila Moreira &#9993</option>
				<option value='transfer&colon;&ast;991639 XML voip01&period;dial&period;opssbank&period;com&period;br' >1639 Gabrielle Moreira &#9993</option>
				<option value='transfer&colon;&ast;991640 XML voip01&period;dial&period;opssbank&period;com&period;br' >1640 Ian &#9993</option>
				<option value='transfer&colon;&ast;991641 XML voip01&period;dial&period;opssbank&period;com&period;br' >1641 Debora Menezes &#9993</option>
				<option value='transfer&colon;&ast;991644 XML voip01&period;dial&period;opssbank&period;com&period;br' >1644 Ana Julia &#9993</option>
				<option value='transfer&colon;&ast;991645 XML voip01&period;dial&period;opssbank&period;com&period;br' >1645 Lorraine Barbosa &#9993</option>
				<option value='transfer&colon;&ast;991646 XML voip01&period;dial&period;opssbank&period;com&period;br' >1646 Natalia Duarte &#9993</option>
				<option value='transfer&colon;&ast;991647 XML voip01&period;dial&period;opssbank&period;com&period;br' >1647 Ronaldo Alex &#9993</option>
				<option value='transfer&colon;&ast;991648 XML voip01&period;dial&period;opssbank&period;com&period;br' >1648 Miriam Oliveira &#9993</option>
				<option value='transfer&colon;&ast;991656 XML voip01&period;dial&period;opssbank&period;com&period;br' >1656 Gleiciane Rodrigues &#9993</option>
				<option value='transfer&colon;&ast;991657 XML voip01&period;dial&period;opssbank&period;com&period;br' >1657 Vivian Fernandes &#9993</option>
				<option value='transfer&colon;&ast;991658 XML voip01&period;dial&period;opssbank&period;com&period;br' >1658 Erilaine&nbsp;Santos &#9993</option>
				<option value='transfer&colon;&ast;991659 XML voip01&period;dial&period;opssbank&period;com&period;br' >1659 Fladielle Santos &#9993</option>
				<option value='transfer&colon;&ast;991660 XML voip01&period;dial&period;opssbank&period;com&period;br' >1660 Ana Neres &#9993</option>
				<option value='transfer&colon;&ast;991667 XML voip01&period;dial&period;opssbank&period;com&period;br' >1667 Joao Victor &#9993</option>
				<option value='transfer&colon;&ast;991668 XML voip01&period;dial&period;opssbank&period;com&period;br' >1668 Maria Eduarda &#9993</option>
				<option value='transfer&colon;&ast;991669 XML voip01&period;dial&period;opssbank&period;com&period;br' >1669 Bruna Antonia &#9993</option>
				<option value='transfer&colon;&ast;991670 XML voip01&period;dial&period;opssbank&period;com&period;br' >1670 Daniel Kesley &#9993</option>
				<option value='transfer&colon;&ast;991671 XML voip01&period;dial&period;opssbank&period;com&period;br' >1671 Cecilia Aparecida &#9993</option>
				<option value='transfer&colon;&ast;991672 XML voip01&period;dial&period;opssbank&period;com&period;br' >1672 Elisangela Alves &#9993</option>
				<option value='transfer&colon;&ast;991673 XML voip01&period;dial&period;opssbank&period;com&period;br' >1673 Nadia Priscila &#9993</option>
				<option value='transfer&colon;&ast;991674 XML voip01&period;dial&period;opssbank&period;com&period;br' >1674 Analice Soares &#9993</option>
				<option value='transfer&colon;&ast;991675 XML voip01&period;dial&period;opssbank&period;com&period;br' >1675 Silvania Almeida &#9993</option>
				<option value='transfer&colon;&ast;999996 XML voip01&period;dial&period;opssbank&period;com&period;br' >9996  &#9993</option>
				<option value='transfer&colon;&ast;999997 XML voip01&period;dial&period;opssbank&period;com&period;br' >9997  &#9993</option>
				<option value='transfer&colon;&ast;999998 XML voip01&period;dial&period;opssbank&period;com&period;br' >9998  &#9993</option>
				<option value='transfer&colon;&ast;999999 XML voip01&period;dial&period;opssbank&period;com&period;br' >9999  &#9993</option>
			</optgroup>
			<optgroup label='Other'>
				<option value='transfer&colon;&ast;98 XML voip01&period;dial&period;opssbank&period;com&period;br' >Check Voicemail</option>
				<option value='transfer&colon;&ast;411 XML voip01&period;dial&period;opssbank&period;com&period;br' >Company Directory</option>
				<option value='hangup&colon;' >Hangup</option>
				<option value='transfer&colon;&ast;732 XML voip01&period;dial&period;opssbank&period;com&period;br' >Record</option>
			</optgroup>
		</select>
	<input type='button' id='btn_select_to_input_queue_timeout_action' class='btn' name='' alt='back' onclick='changeToInputqueue_timeout_action(document.getElementById("queue_timeout_action"));this.style.visibility = "hidden";' value='&#9665;'><br />
	Set the action to perform when the max wait time is reached.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Tier Rules Apply
	</td>
	<td class='vtable' align='left'>
		<select class='formfld' name='queue_tier_rules_apply'>
		<option value='true'>True</option>
		<option value='false' selected='selected' >False</option>
		</select>
	<br />
	Set the tier rule rules apply to true or false.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Tier Rule Wait Second
	</td>
	<td class='vtable' align='left'>
	  <input class='formfld' type='number' name='queue_tier_rule_wait_second' maxlength='255' min='0' step='1' value='30'>
	<br />
	Enter the tier rule wait seconds.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Tier Rule Wait Multiply Level
	</td>
	<td class='vtable' align='left'>
		<select class='formfld' name='queue_tier_rule_wait_multiply_level'>
		<option value='true' selected='selected' >True</option>
		<option value='false'>False</option>
		</select>
	<br />
	Set the tier rule wait multiply level to true or false.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Tier Rule No Agent No Wait
	</td>
	<td class='vtable' align='left'>
		<select class='formfld' name='queue_tier_rule_no_agent_no_wait'>
		<option value='true' selected='selected' >True</option>
		<option value='false'>False</option>
		</select>
	<br />
	Enter the tier rule no agent no wait.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Discard Abandoned After
	</td>
	<td class='vtable' align='left'>
	  <input class='formfld' type='number' name='queue_discard_abandoned_after' maxlength='255' min='0' step='1' value='900'>
	<br />
	The number of seconds before the abandoned call is removed from the queue.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Abandoned Resume Allowed
	</td>
	<td class='vtable' align='left'>
		<select class='formfld' name='queue_abandoned_resume_allowed'>
		<option value='false' selected='selected' >False</option>
		<option value='true'>True</option>
		</select>
	<br />
	A caller who has left the queue can resume their position in the queue by calling again before the abandoned call has been discarded.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Caller ID Name Prefix
	</td>
	<td class='vtable' align='left'>
	  <input class='formfld' type='text' name='queue_cid_prefix' maxlength='255' value=''>
	<br />
	Set a prefix on the caller ID name.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
	  Announce Sound
	</td>
	<td class='vtable' align='left'>
	<script>
	var objs;
	
	function changeToInputqueue_announce_sound(obj){
		tb=document.createElement('INPUT');
		tb.type='text';
		tb.name=obj.name;
		tb.className='formfld';
		tb.setAttribute('id', 'queue_announce_sound');
		tb.setAttribute('style', '');
		tb.value=obj.options[obj.selectedIndex].value;
		document.getElementById('btn_select_to_input_queue_announce_sound').style.visibility = 'hidden';
		tbb=document.createElement('INPUT');
		tbb.setAttribute('class', 'btn');
		tbb.setAttribute('style', 'margin-left: 4px;');
		tbb.type='button';
		tbb.value=$('<div />').html('&#9665;').text();
		tbb.objs=[obj,tb,tbb];
		tbb.onclick=function(){ Replacequeue_announce_sound(this.objs); }
		obj.parentNode.insertBefore(tb,obj);
		obj.parentNode.insertBefore(tbb,obj);
		obj.parentNode.removeChild(obj);
		Replacequeue_announce_sound(this.objs);
	}
	
	function Replacequeue_announce_sound(obj){
		obj[2].parentNode.insertBefore(obj[0],obj[2]);
		obj[0].parentNode.removeChild(obj[1]);
		obj[0].parentNode.removeChild(obj[2]);
		document.getElementById('btn_select_to_input_queue_announce_sound').style.visibility = 'visible';
	}
	</script>
	
	<select name='queue_announce_sound' id='queue_announce_sound' class='formfld'>
		<option></option>
		</select>
	<input type='button' id='btn_select_to_input_queue&lowbar;announce&lowbar;sound' class='btn' name='' alt='back' onclick='changeToInputqueue&lowbar;announce&lowbar;sound(document.getElementById("queue&lowbar;announce&lowbar;sound"));this.style.visibility = "hidden";' value='&#9665;'>	<br />
	A sound to play for a caller at specific intervals, as defined in seconds by the Announce Frequency. Full path to the recording is required.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
	  Announce Frequency
	</td>
	<td class='vtable' align='left'>
	  <input class='formfld' type='number' name='queue_announce_frequency' maxlength='255' min='0' step='1' value=''>
	<br />
	How often should we play the announce sound. Enter a number in seconds
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
	  Exit Key
	</td>
	<td class='vtable' align='left'>
	  <input class='formfld' type='text' name='queue_cc_exit_keys' value=''>
	<br />
	Define a key that can be used to exit the queue.
	</td>
	</tr>
	<tr>
	<td class='vncell' valign='top' align='left' nowrap>
		Description
	</td>
	<td class='vtable' align='left'>
		<input class='formfld' type='text' name='queue_description' maxlength='255' value="">
	<br />
	Enter the description.
	</td>
	</tr>
	</table><br><br><input type='hidden' name='call_center_queue_uuid' value='988348b3-465d-449c-8137-ec3031941c84'>
	<input type='hidden' name='dialplan_uuid' value='dfd61123-d156-42da-a8f7-a448b0f39a8a'>
	<input type='hidden' name='c03bcf3d7d24616c5fefa7412fc0325b666c5a37c356037e8fbe3fa4de2e98c7' value='e7456e5dafc739915a87309243995e3573a63d54b787bcffae4ea396c877569b'>
	</form>
				</div>
				<div id='footer'>
					<span class='footer'>&copy; Copyright 2008 - 2024 <a href='http://www.fusionpbx.com' class='footer' target='_blank'>fusionpbx.com</a> All rights reserved.</span>
				</div>
				</div>
			
	</body>
	</html>
	
		`
}

func callcenterWithoutBranchPageContent() string {
	return `
	<!DOCTYPE html>

<html xmlns='http://www.w3.org/1999/xhtml' xml:lang='en' lang='en'>
<head>
<meta charset='utf-8'>
<meta http-equiv='Content-Type' content='text/html; charset=UTF-8'>
<meta http-equiv='X-UA-Compatible' content='IE=edge'>
<meta name='viewport' content='width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no' />
<meta name="robots" content="noindex, nofollow, noarchive" />

<link rel='stylesheet' type='text/css' href='/resources/bootstrap/css/bootstrap.min.css.php'>
<link rel='stylesheet' type='text/css' href='/resources/bootstrap/css/bootstrap-tempusdominus.min.css.php'>
<link rel='stylesheet' type='text/css' href='/resources/bootstrap/css/bootstrap-colorpicker.min.css.php'>
<link rel='stylesheet' type='text/css' href='/resources/fontawesome/css/all.min.css.php'>
<link rel='stylesheet' type='text/css' href='/themes/default/css.php'>

<link rel='icon' href='/themes/default/favicon.ico'>

<title>Call Center Queue - FusionPBX</title>

<script language='JavaScript' type='text/javascript' src='/resources/jquery/jquery.min.js.php'></script>
<script language='JavaScript' type='text/javascript' src='/resources/jquery/jquery.autosize.input.js.php'></script>
<script language='JavaScript' type='text/javascript' src='/resources/momentjs/moment-with-locales.min.js.php'></script>
<script language='JavaScript' type='text/javascript' src='/resources/bootstrap/js/bootstrap.min.js.php'></script>
<script language='JavaScript' type='text/javascript' src='/resources/bootstrap/js/bootstrap-tempusdominus.min.js.php'></script>
<script language='JavaScript' type='text/javascript' src='/resources/bootstrap/js/bootstrap-colorpicker.min.js.php'></script>
<script language='JavaScript' type='text/javascript' src='/resources/bootstrap/js/bootstrap-pwstrength.min.js.php'></script>
<script language='JavaScript' type='text/javascript'>window.FontAwesomeConfig = { autoReplaceSvg: false }</script>
<script language='JavaScript' type='text/javascript' src='/resources/fontawesome/js/solid.min.js.php' defer></script>


<script language='JavaScript' type='text/javascript'>

//message bar display
	
	function display_message(msg, mood, delay) {
		mood = mood !== undefined ? mood : 'default';
		delay = delay !== undefined ? delay : 1000;
		if (msg !== '') {
			var message_text = $(document.createElement('div'));
			message_text.addClass('message_text message_mood_'+mood);
			message_text.html(msg);
			message_text.on('click', function() {
				var object = $(this);
				object.clearQueue().finish();
				$('#message_container div').remove();
				$('#message_container').css({opacity: 0, 'height': 0}).css({'height': 'auto'});
			} );
			$('#message_container').append(message_text);
			message_text.css({'height': 'auto'}).animate({opacity: 1}, 250, function(){
				$('#message_container').delay(delay).animate({opacity: 0, 'height': 0}, 500, function() {
					$('#message_container div').remove();
					$('#message_container').animate({opacity: 1}, 300).css({'height': 'auto'});
				});
			});
		}
	}
	



$(document).ready(function() {


		//render the messages


//message bar hide on hover
	
	$('#message_container').on('mouseenter',function() {
		$('#message_container div').remove();
		$('#message_container').css({opacity: 0, 'height': 0}).css({'height': 'auto'});
	});
	

//domain selector controls
					
		$('.domain_selector_domain').on('click', function() { show_domains(); });
		$('#header_domain_selector_domain').on('click', function() { show_domains(); });
		$('#domains_hide').on('click', function() { hide_domains(); });

		function show_domains() {
			search_domains('domains_list');

			$('#domains_visible').val(1);
			var scrollbar_width = (window.innerWidth - $(window).width()); //gold: only solution that worked with body { overflow:auto } (add -ms-overflow-style: scrollbar; to <body> style for ie 10+)
			if (scrollbar_width > 0) {
				$('body').css({'margin-right':scrollbar_width, 'overflow':'hidden'}); //disable body scroll bars
				$('.navbar').css('margin-right',scrollbar_width); //adjust navbar margin to compensate
				$('#domains_container').css('right',-scrollbar_width); //domain container right position to compensate
			}
			$(document).scrollTop(0);
			$('#domains_container').show();
			$('#domains_block').animate({marginRight: '+=300'}, 400, function() {
				$('#domains_search').trigger('focus');
			});
		}

		function hide_domains() {
			$('#domains_visible').val(0);
			$(document).ready(function() {
				$('#domains_block').animate({marginRight: '-=300'}, 400, function() {
					$('#domains_search').val('');
					$('.navbar').css('margin-right','0'); //restore navbar margin
					$('#domains_container').css('right','0'); //domain container right position
					$('#domains_container').hide();
					$('body').css({'margin-right':'0','overflow':'auto'}); //enable body scroll bars
					document.activeElement.blur();
				});
			});
		}
		
	
//keyboard shortcut scripts

//key: [enter] - retain default behavior to submit form, when present - note: safari does not honor the first submit element when hiding it using 'display: none;' in the setAttribute method
					
		var action_bar_actions, first_form, first_submit, modal_input_class, modal_continue_button;
		action_bar_actions = document.querySelector('div#action_bar.action_bar > div.actions');
		first_form = document.querySelector('form#frm');

		if (action_bar_actions !== null) {
			if (first_form !== null) {
				first_submit = document.createElement('input');
				first_submit.type = 'submit';
				first_submit.id = 'default_submit';
				first_submit.setAttribute('style','position: absolute; left: -10000px; top: auto; width: 1px; height: 1px; overflow: hidden;');
				first_form.prepend(first_submit);
				window.addEventListener('keydown',function(e){
					modal_input_class = e.target.className;
					if (e.which == 13 && (e.target.tagName == 'INPUT' || e.target.tagName == 'SELECT')) {
						if (modal_input_class.includes('modal-input')) {
							e.preventDefault();
							modal_continue_button = document.getElementById(e.target.dataset.continue);
							if (modal_continue_button) { modal_continue_button.click(); }
						}
						else {
							if (typeof window.submit_form === 'function') { submit_form(); }
							else { document.getElementById('frm').submit(); }
						}
					}
				});
			}
		}
		
	
//common (used by delete and toggle)
					var list_checkboxes;
		list_checkboxes = document.querySelectorAll('table.list tr.list-row td.checkbox input[type=checkbox]');
	
//keyup event listener
	
	window.addEventListener('keyup', function(e) {
		

		//key: [escape] - close modal window, if open, or toggle domain selector
			
			if (e.which == 27) {
				e.preventDefault();
				var modals, modal_visible, modal;
				modal_visible = false;
				modals = document.querySelectorAll('div.modal-window');
				if (modals.length !== 0) {
					for (var x = 0, max = modals.length; x < max; x++) {
						modal = document.getElementById(modals[x].id);
						if (window.getComputedStyle(modal).getPropertyValue('opacity') == 1) {
							modal_visible = true;
						}
					}
				}
				if (modal_visible) {
					modal_close();
				}
				
											
					else {
						if (document.getElementById('domains_visible').value == 0) {
							show_domains();
						}
						else {
							hide_domains();
						}
					}
					
										
			}
			

		//key: [insert], list: to add
									
				if (e.which == 45 && !(e.target.tagName == 'INPUT' && e.target.type == 'text') && e.target.tagName != 'TEXTAREA') {
					e.preventDefault();
					var add_button;
					add_button = document.getElementById('btn_add');
					if (add_button === null || add_button === undefined) {
						add_button = document.querySelector('button[name=btn_add]');
					}
					if (add_button !== null) { add_button.click(); }
				}
				
			
		//key: [delete], list: to delete checked, edit: to delete
			
		//key: [space], list,edit:prevent default space key behavior when opening toggle confirmation (which would automatically *click* the focused continue button on key-up)
									
				if (e.which == 32 && e.target.id == 'btn_toggle') {
					e.preventDefault();
				}
				
			
//keyup end
	
	});
	

//keydown event listener
	
	window.addEventListener('keydown', function(e) {
		

		//key: [space], list: to toggle checked - note: for default [space] checkbox behavior (ie. toggle focused checkbox) include in the if statement: && !(e.target.tagName == 'INPUT' && e.target.type == 'checkbox')
									
				if (e.which == 32 && !(e.target.tagName == 'INPUT' && e.target.type == 'text') && e.target.tagName != 'BUTTON' && !(e.target.tagName == 'INPUT' && e.target.type == 'button') && !(e.target.tagName == 'INPUT' && e.target.type == 'submit') && e.target.tagName != 'TEXTAREA' && list_checkboxes.length !== 0) {
					e.preventDefault();
					var toggle_button;
					toggle_button = document.querySelector('button[name=btn_toggle]');
					if (toggle_button === null || toggle_button === undefined) {
						toggle_button = document.getElementById('btn_toggle');
					}
					if (toggle_button !== null) { toggle_button.click(); }
				}
				
			
		//key: [ctrl]+[a], list,edit: to check all
									
				if ((((e.which == 97 || e.which == 65) && (e.ctrlKey || e.metaKey) && !e.shiftKey) || e.which == 19) && !(e.target.tagName == 'INPUT' && e.target.type == 'text') && e.target.tagName != 'TEXTAREA') {
					var all_checkboxes;
					all_checkboxes = document.querySelectorAll('table.list tr.list-header th.checkbox input[name=checkbox_all]');
					if (typeof all_checkboxes != 'object' || all_checkboxes.length == 0) {
						all_checkboxes = document.querySelectorAll('td.edit_delete_checkbox_all > span > input[name=checkbox_all]');
					}
					if (typeof all_checkboxes == 'object' && all_checkboxes.length > 0) {
						e.preventDefault();
						for (var x = 0, max = all_checkboxes.length; x < max; x++) {
							all_checkboxes[x].click();
						}
					}
				}
				

			
		//key: [ctrl]+[s], edit: to save
									
				if (((e.which == 115 || e.which == 83) && (e.ctrlKey || e.metaKey) && !e.shiftKey) || (e.which == 19)) {
					e.preventDefault();
					var save_button;
					save_button = document.getElementById('btn_save');
					if (save_button === null || save_button === undefined) {
						save_button = document.querySelector('button[name=btn_save]');
					}
					if (save_button !== null) { save_button.click(); }
				}
				
			
		//key: [ctrl]+[c], list,edit: to copy
																
					if (
						(
							(
								(e.which == 99 || e.which == 67) &&
								(e.ctrlKey || e.metaKey) &&
								!e.shiftKey
							) ||
							e.which == 19
						) &&
						!(e.target.tagName == 'INPUT' && e.target.type == 'text') &&
						e.target.tagName != 'TEXTAREA'
						) {
					
										
					var current_selection, copy_button;
					current_selection = window.getSelection();
					if (current_selection === null || current_selection === undefined || current_selection.toString() == '') {
						e.preventDefault();
						copy_button = document.querySelector('button[name=btn_copy]');
						if (copy_button === null || copy_button === undefined) {
							copy_button = document.getElementById('btn_copy');
						}
						if (copy_button !== null) { copy_button.click(); }
					}
				}
				
			
//keydown end
	
	});
	

//link list rows
	
	$('.tr_hover tr,.list tr').each(function(i,e) {
		$(e).children('td:not(.list_control_icon,.list_control_icons,.tr_link_void,.list-row > .no-link,.list-row > .checkbox,.list-row > .button,.list-row > .action-button)').on('click', function() {
			var href = $(this).closest('tr').attr('href');
			var target = $(this).closest('tr').attr('target');
			if (href) {
				if (target) { window.open(href, target); }
				else { window.location = href; }
			}
		});
	});
	

//autosize jquery autosize plugin on applicable input fields
	
	$('input[type=text].txt.auto-size,input[type=number].txt.auto-size,input[type=password].txt.auto-size,input[type=text].formfld.auto-size,input[type=number].formfld.auto-size,input[type=password].formfld.auto-size').autosizeInput();
	

//initialize bootstrap tempusdominus (calendar/datetime picker) plugin
	
	$(function() {
		//set defaults
			$.fn.datetimepicker.Constructor.Default = $.extend({}, $.fn.datetimepicker.Constructor.Default, {
				buttons: {
					showToday: true,
					showClear: true,
					showClose: true,
				},
				icons: {
					time: 'fas fa-clock',
					date: 'fas fa-calendar-alt',
					up: 'fas fa-arrow-up',
					down: 'fas fa-arrow-down',
					previous: 'fas fa-chevron-left',
					next: 'fas fa-chevron-right',
					today: 'fas fa-calendar-check',
					clear: 'fas fa-trash',
					close: 'fas fa-times',
				}
			});
		//define formatting of individual classes
			$('.datepicker').datetimepicker({ format: 'YYYY-MM-DD', });
			$('.datetimepicker').datetimepicker({ format: 'YYYY-MM-DD HH:mm', });
			$('.datetimepicker-future').datetimepicker({ format: 'YYYY-MM-DD HH:mm', minDate: new Date(), });
			$('.datetimesecpicker').datetimepicker({ format: 'YYYY-MM-DD HH:mm:ss', });
	});
	

//apply bootstrap colorpicker plugin
	
	$(function(){
		$('.colorpicker').colorpicker({
			align: 'left',
			customClass: 'colorpicker-2x',
			sliders: {
				saturation: {
					maxLeft: 200,
					maxTop: 200
				},
				hue: {
					maxTop: 200
				},
				alpha: {
					maxTop: 200
				}
			}
		});
	});
	

//apply bootstrap password strength plugin
	
	$('#password').pwstrength({
		common: {
			minChar: 8,
			usernameField: '#username',
		},
		//rules: { },
		ui: {
			colorClasses: ['danger', 'warning', 'warning', 'warning', 'success', 'success'], //weak,poor,normal,medium,good,strong
			progressBarMinPercentage: 15,
			showVerdicts: false,
			viewports: {
				progress: '#pwstrength_progress'
			}
		}
	});
	

//crossfade menu brand images (if hover version set)
	
//generate resizeEnd event after window resize event finishes (used when side menu and on messages app)
	
	$(window).on('resize', function() {
		if (this.resizeTO) { clearTimeout(this.resizeTO); }
		this.resizeTO = setTimeout(function() { $(this).trigger('resizeEnd'); }, 180);
	});
	

//side menu: adjust content container width after window resize
	

}); //document ready end



//audio playback functions

var recording_audio, audio_clock;

function recording_play(recording_id) {
	if (document.getElementById('recording_progress_bar_'+recording_id)) {
		document.getElementById('recording_progress_bar_'+recording_id).style.display='';
	}
	recording_audio = document.getElementById('recording_audio_'+recording_id);

	if (recording_audio.paused) {
		recording_audio.volume = 1;
		recording_audio.play();
		document.getElementById('recording_button_'+recording_id).innerHTML = "<span class='fas fa-pause fa-fw'></span>";
		audio_clock = setInterval(function () { update_progress(recording_id); }, 20);

		$('[id*=recording_button]').not('[id*=recording_button_'+recording_id+']').html("<span class='fas fa-play fa-fw'></span>");
		$('[id*=recording_progress_bar]').not('[id*=recording_progress_bar_'+recording_id+']').css('display', 'none');

		$('audio').each(function(){$('#menu_side_container').width()
			if ($(this).get(0) != recording_audio) {
				$(this).get(0).pause(); //stop playing
				$(this).get(0).currentTime = 0; //reset time
			}
		});
	}
	else {
		recording_audio.pause();
		document.getElementById('recording_button_'+recording_id).innerHTML = "<span class='fas fa-play fa-fw'></span>";
		clearInterval(audio_clock);
	}
}

function recording_stop(recording_id) {
	recording_reset(recording_id);
	clearInterval(audio_clock);
}

function recording_reset(recording_id) {
	recording_audio = document.getElementById('recording_audio_'+recording_id);
	recording_audio.pause();
	recording_audio.currentTime = 0;
	if (document.getElementById('recording_progress_bar_'+recording_id)) {
		document.getElementById('recording_progress_bar_'+recording_id).style.display='none';
	}
	document.getElementById('recording_button_'+recording_id).innerHTML = "<span class='fas fa-play fa-fw'></span>";
	clearInterval(audio_clock);
}

function update_progress(recording_id) {
	recording_audio = document.getElementById('recording_audio_'+recording_id);
	var recording_progress = document.getElementById('recording_progress_'+recording_id);
	var value = 0;
	if (recording_audio.currentTime > 0) {
		value = (100 / recording_audio.duration) * recording_audio.currentTime;
	}
	recording_progress.style.marginLeft = value + '%';
	if (parseInt(recording_audio.duration) > 30) { //seconds
		clearInterval(audio_clock);
	}
}


//handle action bar style on scroll

window.addEventListener('scroll', function(){
	action_bar_scroll('action_bar', 20);
}, false);
function action_bar_scroll(action_bar_id, scroll_position, function_sticky, function_inline) {
	if (document.getElementById(action_bar_id)) {
		//sticky
			if (this.scrollY > scroll_position) {
				document.getElementById(action_bar_id).classList.add('scroll');
				if (typeof function_sticky === 'function') { function_sticky(); }
			}
		//inline
			if (this.scrollY < scroll_position) {
				document.getElementById(action_bar_id).classList.remove('scroll');
				if (typeof function_inline === 'function') { function_inline(); }
			}
	}
}


//enable button class button

function button_enable(button_id) {
	button = document.getElementById(button_id);
	button.disabled = false;
	button.classList.remove('disabled');
	if (button.parentElement.nodeName == 'A') {
		anchor = button.parentElement;
		anchor.classList.remove('disabled');
		anchor.setAttribute('onclick','');
	}
}


//disable button class button

function button_disable(button_id) {
	button = document.getElementById(button_id);
	button.disabled = true;
	button.classList.add('disabled');
	if (button.parentElement.nodeName == 'A') {
		anchor = button.parentElement;
		anchor.classList.add('disabled');
		anchor.setAttribute('onclick','return false;');
	}
}


//checkbox on change

function checkbox_on_change(checkbox) {
	checked = false;
	var inputs = document.getElementsByTagName('input');
	for (var i = 0, max = inputs.length; i < max; i++) {
		if (inputs[i].type === 'checkbox' && inputs[i].checked == true) {
			checked = true;
			break;
		}
	}
	btn_copy = document.getElementById("btn_copy");
	btn_toggle = document.getElementById("btn_toggle");
	btn_delete = document.getElementById("btn_delete");
	if (checked == true) {
		if (btn_copy) {
			btn_copy.style.display = "inline";
		}
		if (btn_toggle) {
			btn_toggle.style.display = "inline";
		}
		if (btn_delete) {
			btn_delete.style.display = "inline";
		}
	}
	 else {
		if (btn_copy) {
			btn_copy.style.display = "none";
		}
		if (btn_toggle) {
			btn_toggle.style.display = "none";
		}
		if (btn_delete) {
			btn_delete.style.display = "none";
		}
	 }
}


//list page functions

function list_all_toggle(modifier) {
	var checkboxes = (modifier !== undefined) ? document.getElementsByClassName('checkbox_'+modifier) : document.querySelectorAll("input[type='checkbox']");
	var checkbox_checked = document.getElementById('checkbox_all' + (modifier !== undefined ? '_'+modifier : '')).checked;
	for (var i = 0, max = checkboxes.length; i < max; i++) {
		checkboxes[i].checked = checkbox_checked;
	}
	if (document.getElementById('btn_check_all') && document.getElementById('btn_check_none')) {
		if (checkbox_checked) {
			document.getElementById('btn_check_all').style.display = 'none';
			document.getElementById('btn_check_none').style.display = '';
		}
		else {
			document.getElementById('btn_check_all').style.display = '';
			document.getElementById('btn_check_none').style.display = 'none';
		}
	}
}

function list_all_check() {
	var inputs = document.getElementsByTagName('input');
	document.getElementById('checkbox_all').checked;
	for (var i = 0, max = inputs.length; i < max; i++) {
		if (inputs[i].type === 'checkbox') {
			inputs[i].checked = true;
		}
	}
}

function list_self_check(checkbox_id) {
	var inputs = document.getElementsByTagName('input');
	for (var i = 0, max = inputs.length; i < max; i++) {
		if (inputs[i].type === 'checkbox' && inputs[i].name.search['enabled'] == -1) {
			inputs[i].checked = false;
		}
	}
	document.getElementById(checkbox_id).checked = true;
}

function list_action_set(action) {
	document.getElementById('action').value = action;
}

function list_form_submit(form_id) {
	document.getElementById(form_id).submit();
}

function list_search_reset() {
	document.getElementById('btn_reset').style.display = 'none';
	document.getElementById('btn_search').style.display = '';
}


//edit page functions

function edit_all_toggle(modifier) {
	var checkboxes = document.getElementsByClassName('checkbox_'+modifier);
	var checkbox_checked = document.getElementById('checkbox_all_'+modifier).checked;
	if (checkboxes.length > 0) {
		for (var i = 0; i < checkboxes.length; ++i) {
			checkboxes[i].checked = checkbox_checked;
		}
		if (document.getElementById('btn_delete')) {
			document.getElementById('btn_delete').value = checkbox_checked ? '' : 'delete';
		}
	}
}

function edit_delete_action(modifier) {
	var checkboxes = document.getElementsByClassName('chk_delete');
	if (document.getElementById('btn_delete') && checkboxes.length > 0) {
		var checkbox_checked = false;
		for (var i = 0; i < checkboxes.length; ++i) {
			if (checkboxes[i].checked) {
				checkbox_checked = true;
			}
			else {
				if (document.getElementById('checkbox_all'+(modifier !== undefined ? '_'+modifier : ''))) {
					document.getElementById('checkbox_all'+(modifier !== undefined ? '_'+modifier : '')).checked = false;
				}
			}
		}
		document.getElementById('btn_delete').value = checkbox_checked ? '' : 'delete';
	}
}


//modal functions

function modal_open(modal_id, focus_id) {
	var modal = document.getElementById(modal_id);
	modal.style.opacity = '1';
	modal.style.pointerEvents = 'auto';
	if (focus_id !== undefined) {
		document.getElementById(focus_id).focus();
	}
}

function modal_close() {
	var modals = document.getElementsByClassName('modal-window');
	if (modals.length > 0) {
		for (var m = 0; m < modals.length; ++m) {
			modals[m].style.opacity = '0';
			modals[m].style.pointerEvents = 'none';
		}
	}
	document.activeElement.blur();
}


//misc functions

function swap_display(a_id, b_id, display_value) {
	display_value = display_value !== undefined ? display_value : 'inline-block';
	a = document.getElementById(a_id);
	b = document.getElementById(b_id);
	if (window.getComputedStyle(a).display === 'none') {
		a.style.display = display_value;
		b.style.display = 'none';
	}
	else {
		a.style.display = 'none';
		b.style.display = display_value;
	}
}

function hide_password_fields() {
	var password_fields = document.querySelectorAll("input[type='password']");
	for (var p = 0, max = password_fields.length; p < max; p++) {
		password_fields[p].style.visibility = 'hidden';
		password_fields[p].type = 'text';
	}
}

window.addEventListener('beforeunload', function(e){
	hide_password_fields();
});


	
function search_domains(element_id) {
var xhttp = new XMLHttpRequest();
xhttp.onreadystatechange = function() {
	//if (this.readyState == 4 && this.status == 200) {
	//	document.getElementById(element_id).innerHTML = this.responseText;
	//}

	//remove current options
	document.getElementById(element_id).innerHTML = '';

	if (this.readyState == 4 && this.status == 200) {

		//create the json object from the response
		obj = JSON.parse(this.responseText);

		//update the domain count
		document.getElementById('domain_count').innerText = '('+ obj.length +')';

		//add new options from the json results
		for (var i=0; i < obj.length; i++) {
			
			//get the variables
			domain_uuid = obj[i].domain_uuid;
			domain_name = obj[i].domain_name;
			if (obj[i].domain_description != null) {
			//	domain_description = DOMPurify.sanitize(obj[i].domain_description);
			}

			//create a div element
			var div = document.createElement('div');

			//add a div title
			div.title = obj[i].domain_name;

			//add a css class
			div.classList.add("domains_list_item");

			//alternate the background color
			if(i%2==0) {
				div.style.background = '#eaedf2';
			}
			else {
				div.style.background = '#ffffff';
			}

			//set the active domain style 
			if ('46ce3c4e-4d9e-420b-9190-116820650089' == obj[i].domain_uuid) {
				div.style.background = '#eeffee';
				div.style.fontWeight = 'bold';
				//div.classList.add("domains_list_item_active");
				//var item_description_class = 'domain_active_list_item_description';
			}
			else {
				//div.classList.add("domains_list_item_inactive");
				//var item_description_class = 'domain_inactive_list_item_description';
			}

			//set link on domain div in list
			div.setAttribute('onclick',"window.location.href='/core/domains/domains.php?domain_uuid=" + obj[i].domain_uuid + "&domain_change=true';");

			//define domain link text and description (if any)
			link_label = obj[i].domain_name;
			if (obj[i].domain_description != null) {
				link_label += " <span class='domain_list_item_description' title=\"" + obj[i].domain_description + "\">" + obj[i].domain_description + "</span>";
			}
			var a_tag = document.createElement('a');
			a_tag.setAttribute('href','manage:'+obj[i].domain_name);
			a_tag.setAttribute('onclick','event.preventDefault();');
			a_tag.innerHTML = link_label;
			div.appendChild(a_tag);

			document.getElementById(element_id).appendChild(div);
		}
	}
};
search = document.getElementById('domains_search');
if (search.value) {
	//xhttp.open("GET", "/core/domains/domain_list.php?search="+search.value, true);
	xhttp.open("GET", "/core/domains/domain_json.php?search="+search.value+"&1bbacd389fd143a75de2ebb27f2a411615b51b6539c7bf7e68100c6dce059c79=0edc8691d57407f6cb47b22dfe01460bf59dc92c3a077c5060c902ad66cb6cca", true);
}
else {
	//xhttp.open("GET", "/core/domains/domain_list.php", true);
	xhttp.open("GET", "/core/domains/domain_json.php?1bbacd389fd143a75de2ebb27f2a411615b51b6539c7bf7e68100c6dce059c79=0edc8691d57407f6cb47b22dfe01460bf59dc92c3a077c5060c902ad66cb6cca", true);
}
xhttp.send();
}
</script>

</head>
<body>

	<div id='message_container'></div>

	
	<div id='domains_container'>
		<input type='hidden' id='domains_visible' value='0'>
		<div id='domains_block'>
			<div id='domains_header'>
				<input id='domains_hide' type='button' class='btn' style='float: right' value="Close">
				<a id='domains_title' href='/core/domains/domains.php'>Domains <span id='domain_count' style='font-size: 80%;'></span></a>
				<br><br>
				<input type='text' id='domains_search' class='formfld' style='margin-left: 0; min-width: 100%; width: 100%;' placeholder="Search..." onkeyup="search_domains('domains_list');">
			</div>
			<div id='domains_list'></div>
		</div>
	</div>


	<div id='qr_code_container' style='display: none;' onclick='$(this).fadeOut(400);'>
	<table cellpadding='0' cellspacing='0' border='0' width='100%' height='100%'><tr><td align='center' valign='middle'>
		<span id='qr_code' onclick="$('#qr_code_container').fadeOut(400);"></span>
	</td></tr></table>
</div>

								 <nav class='navbar navbar-expand-sm fixed-top' >
<div class='container-fluid' style='width: calc(90% - 20px); padding: 0;'>
<div class='navbar-brand'>
	<a href='/'>				<img id='menu_brand_image' class='navbar-logo' src='/themes/default/images/logo.png' title="FusionPBX"></a>
	<a style='margin: 0;'></a>
</div>
<button type='button' class='navbar-toggler' data-toggle='collapse' data-target='#main_navbar' aria-expanded='false' aria-controls='main_navbar' aria-label='Toggle Menu'>
	<span class='fas fa-bars'></span>
</button>
<div class='collapse navbar-collapse' id='main_navbar'>
	<ul class='navbar-nav'>
		<li class='nav-item dropdown '>
			<a class='nav-link' data-toggle='dropdown'  href='#' >
				<span class='fas fa-home' title="Home"></span>
<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Home</span>
			</a>
			<ul class='dropdown-menu'>
				<li class='nav-item'><a class='nav-link' href='/core/users/user_edit.php?id=user' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Account Settings</a></li>
				<li class='nav-item'><a class='nav-link' href='/core/dashboard/' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Dashboard</a></li>
				<li class='nav-item'><a class='nav-link' href='/logout.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Logout</a></li>
			</ul>
		</li>
		<li class='nav-item dropdown '>
			<a class='nav-link' data-toggle='dropdown'  href='#' >
				<span class='fas fa-user' title="Accounts"></span>
<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Accounts</span>
			</a>
			<ul class='dropdown-menu'>
				<li class='nav-item'><a class='nav-link' href='/app/devices/devices.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Devices</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/extensions/extensions.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Extensions</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/gateways/gateways.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Gateways</a></li>
				<li class='nav-item'><a class='nav-link' href='/core/users/users.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Users</a></li>
			</ul>
		</li>
		<li class='nav-item dropdown '>
			<a class='nav-link' data-toggle='dropdown'  href='#' >
				<span class='fas fa-exchange-alt' title="Dialplan"></span>
<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Dialplan</span>
			</a>
			<ul class='dropdown-menu'>
				<li class='nav-item'><a class='nav-link' href='/app/destinations/destinations.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Destinations</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/dialplans/dialplans.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Dialplan Manager</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/dialplans/dialplans.php?app_uuid=c03b422e-13a8-bd1b-e42b-b6b9b4d27ce4' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Inbound Routes</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/dialplans/dialplans.php?app_uuid=8c914ec3-9fc0-8ab5-4cda-6c9288bdc9a3' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Outbound Routes</a></li>
			</ul>
		</li>
		<li class='nav-item dropdown '>
			<a class='nav-link' data-toggle='dropdown'  href='#' >
				<span class='fas fa-paper-plane' title="Applications"></span>
<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Applications</span>
			</a>
			<ul class='dropdown-menu'>
				<li class='nav-item'><a class='nav-link' href='/app/bridges/bridges.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Bridges</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/call_block/call_block.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Block</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/call_broadcast/call_broadcast.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Broadcast</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/call_centers/call_center_queues.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Centers</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/xml_cdr/xml_cdr.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Detail Records</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/call_flows/call_flows.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Flows</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/call_forward/call_forward.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Forward</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/call_recordings/call_recordings.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Call Recordings</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/conference_centers/conference_centers.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Conference Centers</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/conference_controls/conference_controls.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Conference Controls</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/conference_profiles/conference_profiles.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Conference Profiles</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/conferences/conferences.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Conferences</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/contacts/contacts.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Contacts</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/fax/fax.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Fax Server</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/call_forward/call_forward.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Follow Me</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/gswave/index.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>GS Wave</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/ivr_menus/ivr_menus.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>IVR Menus</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/music_on_hold/music_on_hold.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Music on Hold</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/basic_operator_panel/index.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Operator Panel</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/phrases/phrases.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Phrases</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/dialplans/dialplans.php?app_uuid=16589224-c876-aeb3-f59f-523a1c0801f7' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Queues</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/recordings/recordings.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Recordings</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/ring_groups/ring_groups.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Ring Groups</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/streams/streams.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Streams</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/time_conditions/time_conditions.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Time Conditions</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/voicemails/voicemails.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Voicemail</a></li>
			</ul>
		</li>
		<li class='nav-item dropdown '>
			<a class='nav-link' data-toggle='dropdown'  href='#' >
				<span class='fas fa-chart-bar' title="Status"></span>
<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Status</span>
			</a>
			<ul class='dropdown-menu'>
				<li class='nav-item'><a class='nav-link' href='/app/call_center_active/call_center_queue.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Active Call Center</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/calls_active/calls_active.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Active Calls</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/conferences_active/conferences_active.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Active Conferences</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/fifo_list/fifo_list.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Active Queues</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/call_centers/call_center_agent_status.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Agent Status</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/xml_cdr/xml_cdr_statistics.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>CDR Statistics</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/email_queue/email_queue.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Email Queue</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/event_guard/event_guard_logs.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Event Guard</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/xml_cdr/xml_cdr_extension_summary.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Extension Summary</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/fax_queue/fax_queue.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>FAX Queue</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/log_viewer/log_viewer.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Log Viewer</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/registrations/registrations.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Registrations</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/sip_status/sip_status.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>SIP Status</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/system/system.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>System Status</a></li>
				<li class='nav-item'><a class='nav-link' href='/core/user_logs/user_logs.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>User Logs</a></li>
			</ul>
		</li>
		<li class='nav-item dropdown '>
			<a class='nav-link' data-toggle='dropdown'  href='#' >
				<span class='fas fa-cog' title="Advanced"></span>
<span class='d-sm-none d-md-none d-lg-inline' style='margin-left: 5px;'>Advanced</span>
			</a>
			<ul class='dropdown-menu'>
				<li class='nav-item'><a class='nav-link' href='/app/access_controls/access_controls.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Access Controls</a></li>
				<li class='nav-item'><a class='nav-link' href='/core/databases/databases.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Databases</a></li>
				<li class='nav-item'><a class='nav-link' href='/core/default_settings/default_settings.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Default Settings</a></li>
				<li class='nav-item'><a class='nav-link' href='/core/domains/domains.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Domains</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/email_templates/email_templates.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Email Templates</a></li>
				<li class='nav-item'><a class='nav-link' href='/core/groups/groups.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Group Manager</a></li>
				<li class='nav-item'><a class='nav-link' href='/core/menu/menu.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Menu Manager</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/modules/modules.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Modules</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/number_translations/number_translations.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Number Translations</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/sip_profiles/sip_profiles.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>SIP Profiles</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/settings/setting_edit.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Settings</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/database_transactions/database_transactions.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Transactions</a></li>
				<li class='nav-item'><a class='nav-link' href='/core/upgrade/index.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Upgrade</a></li>
				<li class='nav-item'><a class='nav-link' href='/app/vars/vars.php' ><span class='fas fa-bar d-inline-block d-sm-none float-left' style='margin: 4px 10px 0 25px;'></span>Variables</a></li>
			</ul>
		</li>
	</ul>
	<ul class='navbar-nav ml-auto'>
<li class='nav-item'>
	<a class='header_domain' href='#' id='header_domain_selector_domain' title='Domain Selector [ESC]'><i class='fas fa-globe-americas fa-lg fa-fw' style='margin-top: 6px; margin-right: 5px;'></i>voip01&period;dial&period;opssbank&period;com&period;br</a>		</li>
	</ul>
</div>
</div>
</nav>

		<div class='container-fluid' style='padding: 0;' align='center'>

				<div id='main_content'>
		<script>
var Objs;

function changeToInput(obj){
tb=document.createElement('INPUT');
tb.type='text';
tb.name=obj.name;
tb.setAttribute('class', 'formfld');
tb.value=obj.options[obj.selectedIndex].value;
tbb=document.createElement('INPUT');
tbb.setAttribute('class', 'btn');
tbb.setAttribute('style', 'margin-left: 4px;');
tbb.type='button';
tbb.value=$('<div />').html('&#9665;').text();
tbb.objs=[obj,tb,tbb];
tbb.onclick=function(){ Replace(this.objs); }
obj.parentNode.insertBefore(tb,obj);
obj.parentNode.insertBefore(tbb,obj);
obj.parentNode.removeChild(obj);
}

function Replace(obj){
obj[2].parentNode.insertBefore(obj[0],obj[2]);
obj[0].parentNode.removeChild(obj[1]);
obj[0].parentNode.removeChild(obj[2]);
}
</script>

<form name='frm' id='frm' method='post'>
<div class='action_bar' id='action_bar'>
<div class='heading'><b>Call Center Queue</b></div>
<div class='actions'>
<a href='call_center_queues.php' target='_self' style='margin-right: 15px; ' ><button type='button' id='btn_back' alt='Back' title='Back' class='btn btn-default ' ><span class='fas fa-step-backward fa-fw'></span><span class='button-label  pad'>Back</span></button></a><a href='cmd.php?cmd=reload&id=988348b3-465d-449c-8137-ec3031941c84' target='_self' ><button type='button' alt='Reload' title='Reload' class='btn btn-default ' ><span class='fas fa-redo-alt fa-fw'></span><span class='button-label  pad'>Reload</span></button></a><a href='/app/call_center_active/call_center_active.php?queue_name=988348b3-465d-449c-8137-ec3031941c84' target='_self' style='margin-right: 15px; ' ><button type='button' alt='View' title='View' class='btn btn-default ' ><span class='fas fa-eye fa-fw'></span><span class='button-label  pad'>View</span></button></a><button type='submit' id='btn_save' alt='Save' title='Save' class='btn btn-default ' ><span class='fas fa-bolt fa-fw'></span><span class='button-label  pad'>Save</span></button>	</div>
<div style='clear: both;'></div>
</div>
<table width='100%' border='0' cellpadding='0' cellspacing='0'>
<tr>
<td width='30%' class='vncellreq' valign='top' align='left' nowrap>
Queue Name
</td>
<td width='70%' class='vtable' align='left'>
<input class='formfld' type='text' name='queue_name' maxlength='255' value="Teste p&eacute;ricles" required='required'>
<br />
Enter the queue name.
</td>
</tr>
<tr>
<td class='vncellreq' valign='top' align='left' nowrap>
Extension
</td>
<td class='vtable' align='left'>
<input class='formfld' type='number' name='queue_extension' maxlength='255' min='0' step='1' value="200115" required='required'>
<br />
Enter the extension number.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap='nowrap'>
Greeting
</td>
<td class='vtable' align='left'>
<select name='queue_greeting' class='formfld' style='width: 200px;' onchange='changeToInput(this);'>
<option value=''></option>
<optgroup label=Miscellaneous>
<option value='say&colon;'>say</option>
<option value='tone&lowbar;stream&colon;'>tone&lowbar;stream</option>
</optgroup>
</select>
<br />
Select the desired Greeting.
</td>
</tr>
<tr>
<td class='vncellreq' valign='top' align='left' nowrap>
Strategy
</td>
<td class='vtable' align='left'>
<select class='formfld' name='queue_strategy'>
<option value='ring-all'>Ring All</option>
<option value='longest-idle-agent'>Longest Idle Agent</option>
<option value='round-robin' selected='selected'>Round Robin</option>
<option value='top-down'>Top Down</option>
<option value='agent-with-least-talk-time'>Agent With Least Talk Time</option>
<option value='agent-with-fewest-calls'>Agent With Fewest Calls</option>
<option value='sequentially-by-agent-order'>Sequentially By Agent Order</option>
<option value='sequentially-by-next-agent-order'>Sequentially By Next Agent Order</option>
<option value='random'>Random</option>
</select>
<br />
Select the queue ring strategy.
</td>
</tr>
<tr>	<td class='vncell' valign='top'>Agents</td>	<td class='vtable' align='left'>			<table border='0' cellpadding='0' cellspacing='0'>
	<tr>
		<td class='vtable'>Agent Name</td>
		<td class='vtable' style='text-align: center;'>Tier Level</td>
		<td class='vtable' style='text-align: center;'>Tier Position</td>
		<td></td>
	</tr>
<tr>
<td class=''>				<input name='call_center_tiers[0][call_center_tier_uuid]' type='hidden' value="f45b5ff6-043e-4075-83ae-8afe767b5709">
		<select name="call_center_tiers[0][call_center_agent_uuid]" class="formfld" style="width: 200px">
			<option value=""></option>
		<option value="fd7b879d-a2dd-4ef7-9b05-06156a4f80b7">1601</option>
		<option value="6d66dadc-2d38-46e3-ba15-aa113ffa4892">1603</option>
		<option value="c2c0722e-0550-4caa-852e-b765791ef931">1604</option>
		<option value="6a256f0c-84a4-494a-8f66-2913e59779b6">1606</option>
		<option value="a342da12-0d53-4362-8238-417c7a50368c">1607</option>
		<option value="9557f6cf-785d-4e6e-ada6-b94910190556">1609</option>
		<option value="ef50191d-d238-4e44-a29c-db1c22634a21">1610</option>
		<option value="0e96e013-8199-47fd-8268-6c1dcd20aa96">1611</option>
		<option value="d1499462-3461-4d93-9441-72b3c24f857c">1612</option>
		<option value="78352cbe-c627-41cf-a212-a3fc8b8c6260">1613</option>
		<option value="305dbdf8-4546-4053-aa3e-fc268d5f9c91">1614</option>
		<option value="825a6bf5-809f-49fb-a197-7caffa87fd2d">1615</option>
		<option value="98d94ddc-8496-4b95-ae91-9ceee7380083">1616</option>
		<option value="b623066e-8587-4742-8727-a2228c3375d9">1617</option>
		<option value="b6a7e90d-2e94-4967-be9b-48d9c89db12e">1618</option>
		<option value="082eb3b0-fae5-4257-b538-be5e1f9156ca">1619</option>
		<option value="6d2ea436-325d-492c-9143-2b7298ad190a">1626</option>
		<option value="dfabcc60-538e-49dc-8deb-9ab5dedd8c36">1627</option>
		<option value="24d82d98-e94a-4c7f-b64a-feb05261ff6e">1628</option>
		<option value="a6862035-0524-48ee-921c-9b30dc31c2cf">1629</option>
		<option value="3961b4d6-210e-45e9-ac4f-a2ecb66913a5">1630</option>
		<option value="353d6cb2-dbec-44c1-995a-fe5d86e6e277">1631</option>
		<option value="30c7a03d-c175-431d-8c18-d74e1975cd0d">1632</option>
		<option value="fc9a2911-2a09-48ff-b89b-bda47b21998a">1634</option>
		<option value="6fe18ce4-8327-46a8-baab-0f228bcf5649">1635</option>
		<option value="6214f0c8-395c-4c59-87b5-002702195c51">1636</option>
		<option value="f8fba913-a39f-4ed2-94e5-6a51b931f80a">1637</option>
		<option value="7caa54c4-9899-46fb-8177-ebc1f83513d2">1638</option>
		<option value="2e92ca8c-9c72-40d3-9444-04da4bf50a6a">1639</option>
		<option value="f998c21b-f324-47c9-9624-a47876cf0a91">1641</option>
		<option value="d3aa8b8e-54e0-42e9-9f82-65e4fa996796">1642</option>
		<option value="e1f442e1-cbd9-4da9-a9db-2beaf61898b1">1644</option>
		<option value="52b282f8-6499-4b53-a92d-7407c2a601d0">1645</option>
		<option value="50f623e0-3092-428d-bc50-03f086867a82">1646</option>
		<option value="11940fa2-4bd9-433e-b4ad-5b8fb264422d">1647</option>
		<option value="d144d59d-da87-4640-8df5-ab44a4dee372">1649</option>
		<option value="13d3de44-422b-4cf7-aaca-dbb800dfd762">1650</option>
		<option value="864bec58-5173-40b4-8799-5995bb19cb2c">1651</option>
		<option value="6c504577-0c53-4aa5-8da1-c5ca00ebdcd6">1652</option>
		<option value="ac105f6c-6655-46b1-8adf-67e75b3fc6d5">1653</option>
		<option value="264d0e22-a4c8-42e9-947f-a67fab9f2fb9">1654</option>
		<option value="21ee25df-ad7e-4e82-83e7-379670220cdb">1655</option>
		<option value="5baa539e-ab37-4633-9351-42eff7c15edc">1656</option>
		<option value="91f77326-847d-4ed2-8bd1-92ee71180fd3">1657</option>
		<option value="2cdfd4fb-db22-4b31-93b0-febe5dfc61e3">1658</option>
		<option value="598ebf17-4471-475c-8451-7b25d3e1383e">1659</option>
		<option value="dabf7ba9-6d74-48df-b533-dda7219cea07">1660</option>
		<option value="05cf3919-cec9-4b12-9525-b492fa426da3">1661</option>
		<option value="92569c7d-f142-440f-a885-a86c72ad0936">1662</option>
		<option value="015cdc3e-26e7-4347-a2d7-483f0b88ec01">1663</option>
		<option value="1d20eec5-4f2f-43c8-bf2c-d06fb2e272f8">1664</option>
		<option value="97227560-247d-49b7-99c5-c19e3cac9552">1665</option>
		<option value="a958700d-5c31-48d7-9f87-e05cbebfc6c5">1666</option>
		<option value="1c063326-64e2-457c-876d-f536a8ba23a2">1668</option>
		<option value="19deda13-ef7f-43b7-a8c6-3d3f9601ca3c">1669</option>
		<option value="99b094b6-5e08-4472-adc9-853c604be499">1670</option>
		<option value="1265e427-d385-4a3e-a6dc-0093880cd1e5">1671</option>
		<option value="d796a830-cf23-4589-9b07-e610c6d2e7d8">1672</option>
		<option value="dbe0646d-7afb-4ba3-91b0-774952d1021b">1673</option>
		<option value="779905d7-fa83-412f-ad3e-433b53719835">1674</option>
		<option value="b3204faf-86f8-4530-97bd-a6aa22e4886c">1675</option>
		<option value="1bd17098-016f-44ab-b0d0-c6d42fbb944c">1676</option>
		<option value="f81bfa23-068e-41c1-b0b0-643652007a96">1677</option>
		<option value="898d2219-2288-40c6-b2bb-0ecf156cea95">1678</option>
		<option value="5b3d49a1-417d-46b6-b2f2-e6a2f97de078">1679</option>
		<option value="faca8f9b-4f10-403d-9e56-6ab0dae2acbd">1680</option>
		<option value="e9c17f2f-54e8-4a90-852e-fa7d5b1c4f1b">1681</option>
		<option value="9becac36-310b-4852-b464-062e88968c54">1682</option>
		<option value="ce2a697d-85e8-4bd7-a8b5-2f5b7906b883">1683</option>
		<option value="1dde0bf9-966a-40f4-8f02-68928048f0af">1684</option>
		<option value="77dadef4-03b9-4c8a-973e-2b912789a972">1685</option>
		<option value="d93b5abb-5dcf-4e40-af16-7338f5f64c3a">1686</option>
		<option value="92374956-dd15-4edb-a340-70139e0e941c">1687</option>
		<option value="e3923b66-6cf3-4da6-b1aa-caba339127c0">1688</option>
		<option value="567ae4fc-db87-4a4b-b405-265c50d55604">1689</option>
		<option value="7bd3f9e8-b321-4e92-a862-09d416ae1594">1690</option>
		<option value="38ad6832-537f-4f7b-af37-5b8f0f755359">1691</option>
		<option value="ac68ed3a-cf71-4d66-9577-3f6512fed3d1">1692</option>
		<option value="b7ad4660-4bd7-42a0-9223-f96bceac18b6">1693</option>
		<option value="9d3f1651-e68e-44b8-b425-2f697a0109af">1694</option>
		<option value="b70815c9-1672-414a-ae4c-25b39f54beff">1696</option>
		<option value="77cdd4d6-c15b-45c8-ad6c-671d8f9426b6">9998</option>
		<option value="b4ca43fc-d988-46d3-9cc6-2a8ecc5a0893">9999</option>
		</select>		</td>
<td class='' style='text-align: center;'>				 <select name="call_center_tiers[0][tier_level]" class="formfld">
		<option value="0" >0</option>
		<option value="1" >1</option>
		<option value="2" >2</option>
		<option value="3" >3</option>
		<option value="4" >4</option>
		<option value="5" >5</option>
		<option value="6" >6</option>
		<option value="7" >7</option>
		<option value="8" >8</option>
		<option value="9" >9</option>
		</select>
</td>
<td class='' style='text-align: center;'>
		<select name="call_center_tiers[0][tier_position]" class="formfld">
		<option value="0" >0</option>
		<option value="1" >1</option>
		<option value="2" >2</option>
		<option value="3" >3</option>
		<option value="4" >4</option>
		<option value="5" >5</option>
		<option value="6" >6</option>
		<option value="7" >7</option>
		<option value="8" >8</option>
		<option value="9" >9</option>
		</select>
</td>
<td class=''>			<a href="call_center_queue_edit.php?id=988348b3-465d-449c-8137-ec3031941c84&call_center_tier_uuid=f45b5ff6-043e-4075-83ae-8afe767b5709&a=delete" alt="Delete" onclick="return confirm('Do you really want to DELETE this?');"><button type='button' class='btn btn-default list_control_icon'><span class='fas fa-minus'></span></button></a>		</td>
</tr>
<tr>
<td class=''>				<input name='call_center_tiers[1][call_center_tier_uuid]' type='hidden' value="142666e5-66ca-44f5-a05e-5c95fd47ce02">
		<select name="call_center_tiers[1][call_center_agent_uuid]" class="formfld" style="width: 200px">
			<option value=""></option>
		<option value="fd7b879d-a2dd-4ef7-9b05-06156a4f80b7">1601</option>
		<option value="6d66dadc-2d38-46e3-ba15-aa113ffa4892">1603</option>
		<option value="c2c0722e-0550-4caa-852e-b765791ef931">1604</option>
		<option value="6a256f0c-84a4-494a-8f66-2913e59779b6">1606</option>
		<option value="a342da12-0d53-4362-8238-417c7a50368c">1607</option>
		<option value="9557f6cf-785d-4e6e-ada6-b94910190556">1609</option>
		<option value="ef50191d-d238-4e44-a29c-db1c22634a21">1610</option>
		<option value="0e96e013-8199-47fd-8268-6c1dcd20aa96">1611</option>
		<option value="d1499462-3461-4d93-9441-72b3c24f857c">1612</option>
		<option value="78352cbe-c627-41cf-a212-a3fc8b8c6260">1613</option>
		<option value="305dbdf8-4546-4053-aa3e-fc268d5f9c91">1614</option>
		<option value="825a6bf5-809f-49fb-a197-7caffa87fd2d">1615</option>
		<option value="98d94ddc-8496-4b95-ae91-9ceee7380083">1616</option>
		<option value="b623066e-8587-4742-8727-a2228c3375d9">1617</option>
		<option value="b6a7e90d-2e94-4967-be9b-48d9c89db12e">1618</option>
		<option value="082eb3b0-fae5-4257-b538-be5e1f9156ca">1619</option>
		<option value="6d2ea436-325d-492c-9143-2b7298ad190a">1626</option>
		<option value="dfabcc60-538e-49dc-8deb-9ab5dedd8c36">1627</option>
		<option value="24d82d98-e94a-4c7f-b64a-feb05261ff6e">1628</option>
		<option value="a6862035-0524-48ee-921c-9b30dc31c2cf">1629</option>
		<option value="3961b4d6-210e-45e9-ac4f-a2ecb66913a5">1630</option>
		<option value="353d6cb2-dbec-44c1-995a-fe5d86e6e277">1631</option>
		<option value="30c7a03d-c175-431d-8c18-d74e1975cd0d">1632</option>
		<option value="fc9a2911-2a09-48ff-b89b-bda47b21998a">1634</option>
		<option value="6fe18ce4-8327-46a8-baab-0f228bcf5649">1635</option>
		<option value="6214f0c8-395c-4c59-87b5-002702195c51">1636</option>
		<option value="f8fba913-a39f-4ed2-94e5-6a51b931f80a">1637</option>
		<option value="7caa54c4-9899-46fb-8177-ebc1f83513d2">1638</option>
		<option value="2e92ca8c-9c72-40d3-9444-04da4bf50a6a">1639</option>
		<option value="f998c21b-f324-47c9-9624-a47876cf0a91">1641</option>
		<option value="d3aa8b8e-54e0-42e9-9f82-65e4fa996796">1642</option>
		<option value="e1f442e1-cbd9-4da9-a9db-2beaf61898b1">1644</option>
		<option value="52b282f8-6499-4b53-a92d-7407c2a601d0">1645</option>
		<option value="50f623e0-3092-428d-bc50-03f086867a82">1646</option>
		<option value="11940fa2-4bd9-433e-b4ad-5b8fb264422d">1647</option>
		<option value="d144d59d-da87-4640-8df5-ab44a4dee372">1649</option>
		<option value="13d3de44-422b-4cf7-aaca-dbb800dfd762">1650</option>
		<option value="864bec58-5173-40b4-8799-5995bb19cb2c">1651</option>
		<option value="6c504577-0c53-4aa5-8da1-c5ca00ebdcd6">1652</option>
		<option value="ac105f6c-6655-46b1-8adf-67e75b3fc6d5">1653</option>
		<option value="264d0e22-a4c8-42e9-947f-a67fab9f2fb9">1654</option>
		<option value="21ee25df-ad7e-4e82-83e7-379670220cdb">1655</option>
		<option value="5baa539e-ab37-4633-9351-42eff7c15edc">1656</option>
		<option value="91f77326-847d-4ed2-8bd1-92ee71180fd3">1657</option>
		<option value="2cdfd4fb-db22-4b31-93b0-febe5dfc61e3">1658</option>
		<option value="598ebf17-4471-475c-8451-7b25d3e1383e">1659</option>
		<option value="dabf7ba9-6d74-48df-b533-dda7219cea07">1660</option>
		<option value="05cf3919-cec9-4b12-9525-b492fa426da3">1661</option>
		<option value="92569c7d-f142-440f-a885-a86c72ad0936">1662</option>
		<option value="015cdc3e-26e7-4347-a2d7-483f0b88ec01">1663</option>
		<option value="1d20eec5-4f2f-43c8-bf2c-d06fb2e272f8">1664</option>
		<option value="97227560-247d-49b7-99c5-c19e3cac9552">1665</option>
		<option value="a958700d-5c31-48d7-9f87-e05cbebfc6c5">1666</option>
		<option value="1c063326-64e2-457c-876d-f536a8ba23a2">1668</option>
		<option value="19deda13-ef7f-43b7-a8c6-3d3f9601ca3c">1669</option>
		<option value="99b094b6-5e08-4472-adc9-853c604be499">1670</option>
		<option value="1265e427-d385-4a3e-a6dc-0093880cd1e5">1671</option>
		<option value="d796a830-cf23-4589-9b07-e610c6d2e7d8">1672</option>
		<option value="dbe0646d-7afb-4ba3-91b0-774952d1021b">1673</option>
		<option value="779905d7-fa83-412f-ad3e-433b53719835">1674</option>
		<option value="b3204faf-86f8-4530-97bd-a6aa22e4886c">1675</option>
		<option value="1bd17098-016f-44ab-b0d0-c6d42fbb944c">1676</option>
		<option value="f81bfa23-068e-41c1-b0b0-643652007a96">1677</option>
		<option value="898d2219-2288-40c6-b2bb-0ecf156cea95">1678</option>
		<option value="5b3d49a1-417d-46b6-b2f2-e6a2f97de078">1679</option>
		<option value="faca8f9b-4f10-403d-9e56-6ab0dae2acbd">1680</option>
		<option value="e9c17f2f-54e8-4a90-852e-fa7d5b1c4f1b">1681</option>
		<option value="9becac36-310b-4852-b464-062e88968c54">1682</option>
		<option value="ce2a697d-85e8-4bd7-a8b5-2f5b7906b883">1683</option>
		<option value="1dde0bf9-966a-40f4-8f02-68928048f0af">1684</option>
		<option value="77dadef4-03b9-4c8a-973e-2b912789a972">1685</option>
		<option value="d93b5abb-5dcf-4e40-af16-7338f5f64c3a">1686</option>
		<option value="92374956-dd15-4edb-a340-70139e0e941c">1687</option>
		<option value="e3923b66-6cf3-4da6-b1aa-caba339127c0">1688</option>
		<option value="567ae4fc-db87-4a4b-b405-265c50d55604">1689</option>
		<option value="7bd3f9e8-b321-4e92-a862-09d416ae1594">1690</option>
		<option value="38ad6832-537f-4f7b-af37-5b8f0f755359">1691</option>
		<option value="ac68ed3a-cf71-4d66-9577-3f6512fed3d1">1692</option>
		<option value="b7ad4660-4bd7-42a0-9223-f96bceac18b6">1693</option>
		<option value="9d3f1651-e68e-44b8-b425-2f697a0109af">1694</option>
		<option value="b70815c9-1672-414a-ae4c-25b39f54beff">1696</option>
		<option value="77cdd4d6-c15b-45c8-ad6c-671d8f9426b6">9998</option>
		<option value="b4ca43fc-d988-46d3-9cc6-2a8ecc5a0893">9999</option>
		</select>		</td>
<td class='' style='text-align: center;'>				 <select name="call_center_tiers[1][tier_level]" class="formfld">
		<option value="0" >0</option>
		<option value="1" >1</option>
		<option value="2" >2</option>
		<option value="3" >3</option>
		<option value="4" >4</option>
		<option value="5" >5</option>
		<option value="6" >6</option>
		<option value="7" >7</option>
		<option value="8" >8</option>
		<option value="9" >9</option>
		</select>
</td>
<td class='' style='text-align: center;'>
		<select name="call_center_tiers[1][tier_position]" class="formfld">
		<option value="0" >0</option>
		<option value="1" >1</option>
		<option value="2" >2</option>
		<option value="3" >3</option>
		<option value="4" >4</option>
		<option value="5" >5</option>
		<option value="6" >6</option>
		<option value="7" >7</option>
		<option value="8" >8</option>
		<option value="9" >9</option>
		</select>
</td>
<td class=''>			<a href="call_center_queue_edit.php?id=988348b3-465d-449c-8137-ec3031941c84&call_center_tier_uuid=142666e5-66ca-44f5-a05e-5c95fd47ce02&a=delete" alt="Delete" onclick="return confirm('Do you really want to DELETE this?');"><button type='button' class='btn btn-default list_control_icon'><span class='fas fa-minus'></span></button></a>		</td>
</tr>
<tr>
<td class=''>				<input name='call_center_tiers[2][call_center_tier_uuid]' type='hidden' value="5eb0ec8b-8423-4004-890e-206b80caf6bc">
		<select name="call_center_tiers[2][call_center_agent_uuid]" class="formfld" style="width: 200px">
			<option value=""></option>
		<option value="fd7b879d-a2dd-4ef7-9b05-06156a4f80b7">1601</option>
		<option value="6d66dadc-2d38-46e3-ba15-aa113ffa4892">1603</option>
		<option value="c2c0722e-0550-4caa-852e-b765791ef931">1604</option>
		<option value="6a256f0c-84a4-494a-8f66-2913e59779b6">1606</option>
		<option value="a342da12-0d53-4362-8238-417c7a50368c">1607</option>
		<option value="9557f6cf-785d-4e6e-ada6-b94910190556">1609</option>
		<option value="ef50191d-d238-4e44-a29c-db1c22634a21">1610</option>
		<option value="0e96e013-8199-47fd-8268-6c1dcd20aa96">1611</option>
		<option value="d1499462-3461-4d93-9441-72b3c24f857c">1612</option>
		<option value="78352cbe-c627-41cf-a212-a3fc8b8c6260">1613</option>
		<option value="305dbdf8-4546-4053-aa3e-fc268d5f9c91">1614</option>
		<option value="825a6bf5-809f-49fb-a197-7caffa87fd2d">1615</option>
		<option value="98d94ddc-8496-4b95-ae91-9ceee7380083">1616</option>
		<option value="b623066e-8587-4742-8727-a2228c3375d9">1617</option>
		<option value="b6a7e90d-2e94-4967-be9b-48d9c89db12e">1618</option>
		<option value="082eb3b0-fae5-4257-b538-be5e1f9156ca">1619</option>
		<option value="6d2ea436-325d-492c-9143-2b7298ad190a">1626</option>
		<option value="dfabcc60-538e-49dc-8deb-9ab5dedd8c36">1627</option>
		<option value="24d82d98-e94a-4c7f-b64a-feb05261ff6e">1628</option>
		<option value="a6862035-0524-48ee-921c-9b30dc31c2cf">1629</option>
		<option value="3961b4d6-210e-45e9-ac4f-a2ecb66913a5">1630</option>
		<option value="353d6cb2-dbec-44c1-995a-fe5d86e6e277">1631</option>
		<option value="30c7a03d-c175-431d-8c18-d74e1975cd0d">1632</option>
		<option value="fc9a2911-2a09-48ff-b89b-bda47b21998a">1634</option>
		<option value="6fe18ce4-8327-46a8-baab-0f228bcf5649">1635</option>
		<option value="6214f0c8-395c-4c59-87b5-002702195c51">1636</option>
		<option value="f8fba913-a39f-4ed2-94e5-6a51b931f80a">1637</option>
		<option value="7caa54c4-9899-46fb-8177-ebc1f83513d2">1638</option>
		<option value="2e92ca8c-9c72-40d3-9444-04da4bf50a6a">1639</option>
		<option value="f998c21b-f324-47c9-9624-a47876cf0a91">1641</option>
		<option value="d3aa8b8e-54e0-42e9-9f82-65e4fa996796">1642</option>
		<option value="e1f442e1-cbd9-4da9-a9db-2beaf61898b1">1644</option>
		<option value="52b282f8-6499-4b53-a92d-7407c2a601d0">1645</option>
		<option value="50f623e0-3092-428d-bc50-03f086867a82">1646</option>
		<option value="11940fa2-4bd9-433e-b4ad-5b8fb264422d">1647</option>
		<option value="d144d59d-da87-4640-8df5-ab44a4dee372">1649</option>
		<option value="13d3de44-422b-4cf7-aaca-dbb800dfd762">1650</option>
		<option value="864bec58-5173-40b4-8799-5995bb19cb2c">1651</option>
		<option value="6c504577-0c53-4aa5-8da1-c5ca00ebdcd6">1652</option>
		<option value="ac105f6c-6655-46b1-8adf-67e75b3fc6d5">1653</option>
		<option value="264d0e22-a4c8-42e9-947f-a67fab9f2fb9">1654</option>
		<option value="21ee25df-ad7e-4e82-83e7-379670220cdb">1655</option>
		<option value="5baa539e-ab37-4633-9351-42eff7c15edc">1656</option>
		<option value="91f77326-847d-4ed2-8bd1-92ee71180fd3">1657</option>
		<option value="2cdfd4fb-db22-4b31-93b0-febe5dfc61e3">1658</option>
		<option value="598ebf17-4471-475c-8451-7b25d3e1383e">1659</option>
		<option value="dabf7ba9-6d74-48df-b533-dda7219cea07">1660</option>
		<option value="05cf3919-cec9-4b12-9525-b492fa426da3">1661</option>
		<option value="92569c7d-f142-440f-a885-a86c72ad0936">1662</option>
		<option value="015cdc3e-26e7-4347-a2d7-483f0b88ec01">1663</option>
		<option value="1d20eec5-4f2f-43c8-bf2c-d06fb2e272f8">1664</option>
		<option value="97227560-247d-49b7-99c5-c19e3cac9552">1665</option>
		<option value="a958700d-5c31-48d7-9f87-e05cbebfc6c5">1666</option>
		<option value="1c063326-64e2-457c-876d-f536a8ba23a2">1668</option>
		<option value="19deda13-ef7f-43b7-a8c6-3d3f9601ca3c">1669</option>
		<option value="99b094b6-5e08-4472-adc9-853c604be499">1670</option>
		<option value="1265e427-d385-4a3e-a6dc-0093880cd1e5">1671</option>
		<option value="d796a830-cf23-4589-9b07-e610c6d2e7d8">1672</option>
		<option value="dbe0646d-7afb-4ba3-91b0-774952d1021b">1673</option>
		<option value="779905d7-fa83-412f-ad3e-433b53719835">1674</option>
		<option value="b3204faf-86f8-4530-97bd-a6aa22e4886c">1675</option>
		<option value="1bd17098-016f-44ab-b0d0-c6d42fbb944c">1676</option>
		<option value="f81bfa23-068e-41c1-b0b0-643652007a96">1677</option>
		<option value="898d2219-2288-40c6-b2bb-0ecf156cea95">1678</option>
		<option value="5b3d49a1-417d-46b6-b2f2-e6a2f97de078">1679</option>
		<option value="faca8f9b-4f10-403d-9e56-6ab0dae2acbd">1680</option>
		<option value="e9c17f2f-54e8-4a90-852e-fa7d5b1c4f1b">1681</option>
		<option value="9becac36-310b-4852-b464-062e88968c54">1682</option>
		<option value="ce2a697d-85e8-4bd7-a8b5-2f5b7906b883">1683</option>
		<option value="1dde0bf9-966a-40f4-8f02-68928048f0af">1684</option>
		<option value="77dadef4-03b9-4c8a-973e-2b912789a972">1685</option>
		<option value="d93b5abb-5dcf-4e40-af16-7338f5f64c3a">1686</option>
		<option value="92374956-dd15-4edb-a340-70139e0e941c">1687</option>
		<option value="e3923b66-6cf3-4da6-b1aa-caba339127c0">1688</option>
		<option value="567ae4fc-db87-4a4b-b405-265c50d55604">1689</option>
		<option value="7bd3f9e8-b321-4e92-a862-09d416ae1594">1690</option>
		<option value="38ad6832-537f-4f7b-af37-5b8f0f755359">1691</option>
		<option value="ac68ed3a-cf71-4d66-9577-3f6512fed3d1">1692</option>
		<option value="b7ad4660-4bd7-42a0-9223-f96bceac18b6">1693</option>
		<option value="9d3f1651-e68e-44b8-b425-2f697a0109af">1694</option>
		<option value="b70815c9-1672-414a-ae4c-25b39f54beff">1696</option>
		<option value="77cdd4d6-c15b-45c8-ad6c-671d8f9426b6">9998</option>
		<option value="b4ca43fc-d988-46d3-9cc6-2a8ecc5a0893">9999</option>
		</select>		</td>
<td class='' style='text-align: center;'>				 <select name="call_center_tiers[2][tier_level]" class="formfld">
		<option value="0" >0</option>
		<option value="1" >1</option>
		<option value="2" >2</option>
		<option value="3" >3</option>
		<option value="4" >4</option>
		<option value="5" >5</option>
		<option value="6" >6</option>
		<option value="7" >7</option>
		<option value="8" >8</option>
		<option value="9" >9</option>
		</select>
</td>
<td class='' style='text-align: center;'>
		<select name="call_center_tiers[2][tier_position]" class="formfld">
		<option value="0" >0</option>
		<option value="1" >1</option>
		<option value="2" >2</option>
		<option value="3" >3</option>
		<option value="4" >4</option>
		<option value="5" >5</option>
		<option value="6" >6</option>
		<option value="7" >7</option>
		<option value="8" >8</option>
		<option value="9" >9</option>
		</select>
</td>
<td class=''>			<a href="call_center_queue_edit.php?id=988348b3-465d-449c-8137-ec3031941c84&call_center_tier_uuid=5eb0ec8b-8423-4004-890e-206b80caf6bc&a=delete" alt="Delete" onclick="return confirm('Do you really want to DELETE this?');"><button type='button' class='btn btn-default list_control_icon'><span class='fas fa-minus'></span></button></a>		</td>
</tr>
<tr>
<td class=''>				<input name='call_center_tiers[3][call_center_tier_uuid]' type='hidden' value="5ce52dc0-97c1-4e27-9935-dab10d539a92">
		<select name="call_center_tiers[3][call_center_agent_uuid]" class="formfld" style="width: 200px">
			<option value=""></option>
		<option value="fd7b879d-a2dd-4ef7-9b05-06156a4f80b7">1601</option>
		<option value="6d66dadc-2d38-46e3-ba15-aa113ffa4892">1603</option>
		<option value="c2c0722e-0550-4caa-852e-b765791ef931">1604</option>
		<option value="6a256f0c-84a4-494a-8f66-2913e59779b6">1606</option>
		<option value="a342da12-0d53-4362-8238-417c7a50368c">1607</option>
		<option value="9557f6cf-785d-4e6e-ada6-b94910190556">1609</option>
		<option value="ef50191d-d238-4e44-a29c-db1c22634a21">1610</option>
		<option value="0e96e013-8199-47fd-8268-6c1dcd20aa96">1611</option>
		<option value="d1499462-3461-4d93-9441-72b3c24f857c">1612</option>
		<option value="78352cbe-c627-41cf-a212-a3fc8b8c6260">1613</option>
		<option value="305dbdf8-4546-4053-aa3e-fc268d5f9c91">1614</option>
		<option value="825a6bf5-809f-49fb-a197-7caffa87fd2d">1615</option>
		<option value="98d94ddc-8496-4b95-ae91-9ceee7380083">1616</option>
		<option value="b623066e-8587-4742-8727-a2228c3375d9">1617</option>
		<option value="b6a7e90d-2e94-4967-be9b-48d9c89db12e">1618</option>
		<option value="082eb3b0-fae5-4257-b538-be5e1f9156ca">1619</option>
		<option value="6d2ea436-325d-492c-9143-2b7298ad190a">1626</option>
		<option value="dfabcc60-538e-49dc-8deb-9ab5dedd8c36">1627</option>
		<option value="24d82d98-e94a-4c7f-b64a-feb05261ff6e">1628</option>
		<option value="a6862035-0524-48ee-921c-9b30dc31c2cf">1629</option>
		<option value="3961b4d6-210e-45e9-ac4f-a2ecb66913a5">1630</option>
		<option value="353d6cb2-dbec-44c1-995a-fe5d86e6e277">1631</option>
		<option value="30c7a03d-c175-431d-8c18-d74e1975cd0d">1632</option>
		<option value="fc9a2911-2a09-48ff-b89b-bda47b21998a">1634</option>
		<option value="6fe18ce4-8327-46a8-baab-0f228bcf5649">1635</option>
		<option value="6214f0c8-395c-4c59-87b5-002702195c51">1636</option>
		<option value="f8fba913-a39f-4ed2-94e5-6a51b931f80a">1637</option>
		<option value="7caa54c4-9899-46fb-8177-ebc1f83513d2">1638</option>
		<option value="2e92ca8c-9c72-40d3-9444-04da4bf50a6a">1639</option>
		<option value="f998c21b-f324-47c9-9624-a47876cf0a91">1641</option>
		<option value="d3aa8b8e-54e0-42e9-9f82-65e4fa996796">1642</option>
		<option value="e1f442e1-cbd9-4da9-a9db-2beaf61898b1">1644</option>
		<option value="52b282f8-6499-4b53-a92d-7407c2a601d0">1645</option>
		<option value="50f623e0-3092-428d-bc50-03f086867a82">1646</option>
		<option value="11940fa2-4bd9-433e-b4ad-5b8fb264422d">1647</option>
		<option value="d144d59d-da87-4640-8df5-ab44a4dee372">1649</option>
		<option value="13d3de44-422b-4cf7-aaca-dbb800dfd762">1650</option>
		<option value="864bec58-5173-40b4-8799-5995bb19cb2c">1651</option>
		<option value="6c504577-0c53-4aa5-8da1-c5ca00ebdcd6">1652</option>
		<option value="ac105f6c-6655-46b1-8adf-67e75b3fc6d5">1653</option>
		<option value="264d0e22-a4c8-42e9-947f-a67fab9f2fb9">1654</option>
		<option value="21ee25df-ad7e-4e82-83e7-379670220cdb">1655</option>
		<option value="5baa539e-ab37-4633-9351-42eff7c15edc">1656</option>
		<option value="91f77326-847d-4ed2-8bd1-92ee71180fd3">1657</option>
		<option value="2cdfd4fb-db22-4b31-93b0-febe5dfc61e3">1658</option>
		<option value="598ebf17-4471-475c-8451-7b25d3e1383e">1659</option>
		<option value="dabf7ba9-6d74-48df-b533-dda7219cea07">1660</option>
		<option value="05cf3919-cec9-4b12-9525-b492fa426da3">1661</option>
		<option value="92569c7d-f142-440f-a885-a86c72ad0936">1662</option>
		<option value="015cdc3e-26e7-4347-a2d7-483f0b88ec01">1663</option>
		<option value="1d20eec5-4f2f-43c8-bf2c-d06fb2e272f8">1664</option>
		<option value="97227560-247d-49b7-99c5-c19e3cac9552">1665</option>
		<option value="a958700d-5c31-48d7-9f87-e05cbebfc6c5">1666</option>
		<option value="1c063326-64e2-457c-876d-f536a8ba23a2">1668</option>
		<option value="19deda13-ef7f-43b7-a8c6-3d3f9601ca3c">1669</option>
		<option value="99b094b6-5e08-4472-adc9-853c604be499">1670</option>
		<option value="1265e427-d385-4a3e-a6dc-0093880cd1e5">1671</option>
		<option value="d796a830-cf23-4589-9b07-e610c6d2e7d8">1672</option>
		<option value="dbe0646d-7afb-4ba3-91b0-774952d1021b">1673</option>
		<option value="779905d7-fa83-412f-ad3e-433b53719835">1674</option>
		<option value="b3204faf-86f8-4530-97bd-a6aa22e4886c">1675</option>
		<option value="1bd17098-016f-44ab-b0d0-c6d42fbb944c">1676</option>
		<option value="f81bfa23-068e-41c1-b0b0-643652007a96">1677</option>
		<option value="898d2219-2288-40c6-b2bb-0ecf156cea95">1678</option>
		<option value="5b3d49a1-417d-46b6-b2f2-e6a2f97de078">1679</option>
		<option value="faca8f9b-4f10-403d-9e56-6ab0dae2acbd">1680</option>
		<option value="e9c17f2f-54e8-4a90-852e-fa7d5b1c4f1b">1681</option>
		<option value="9becac36-310b-4852-b464-062e88968c54">1682</option>
		<option value="ce2a697d-85e8-4bd7-a8b5-2f5b7906b883">1683</option>
		<option value="1dde0bf9-966a-40f4-8f02-68928048f0af">1684</option>
		<option value="77dadef4-03b9-4c8a-973e-2b912789a972">1685</option>
		<option value="d93b5abb-5dcf-4e40-af16-7338f5f64c3a">1686</option>
		<option value="92374956-dd15-4edb-a340-70139e0e941c">1687</option>
		<option value="e3923b66-6cf3-4da6-b1aa-caba339127c0">1688</option>
		<option value="567ae4fc-db87-4a4b-b405-265c50d55604">1689</option>
		<option value="7bd3f9e8-b321-4e92-a862-09d416ae1594">1690</option>
		<option value="38ad6832-537f-4f7b-af37-5b8f0f755359">1691</option>
		<option value="ac68ed3a-cf71-4d66-9577-3f6512fed3d1">1692</option>
		<option value="b7ad4660-4bd7-42a0-9223-f96bceac18b6">1693</option>
		<option value="9d3f1651-e68e-44b8-b425-2f697a0109af">1694</option>
		<option value="b70815c9-1672-414a-ae4c-25b39f54beff">1696</option>
		<option value="77cdd4d6-c15b-45c8-ad6c-671d8f9426b6">9998</option>
		<option value="b4ca43fc-d988-46d3-9cc6-2a8ecc5a0893">9999</option>
		</select>		</td>
<td class='' style='text-align: center;'>				 <select name="call_center_tiers[3][tier_level]" class="formfld">
		<option value="0" >0</option>
		<option value="1" >1</option>
		<option value="2" >2</option>
		<option value="3" >3</option>
		<option value="4" >4</option>
		<option value="5" >5</option>
		<option value="6" >6</option>
		<option value="7" >7</option>
		<option value="8" >8</option>
		<option value="9" >9</option>
		</select>
</td>
<td class='' style='text-align: center;'>
		<select name="call_center_tiers[3][tier_position]" class="formfld">
		<option value="0" >0</option>
		<option value="1" >1</option>
		<option value="2" >2</option>
		<option value="3" >3</option>
		<option value="4" >4</option>
		<option value="5" >5</option>
		<option value="6" >6</option>
		<option value="7" >7</option>
		<option value="8" >8</option>
		<option value="9" >9</option>
		</select>
</td>
<td class=''>			<a href="call_center_queue_edit.php?id=988348b3-465d-449c-8137-ec3031941c84&call_center_tier_uuid=5ce52dc0-97c1-4e27-9935-dab10d539a92&a=delete" alt="Delete" onclick="return confirm('Do you really want to DELETE this?');"><button type='button' class='btn btn-default list_control_icon'><span class='fas fa-minus'></span></button></a>		</td>
</tr>
<tr>
<td class=''>				<input name='call_center_tiers[4][call_center_tier_uuid]' type='hidden' value="d3f3f74b-328c-485e-a751-9f6999567c95">
		<select name="call_center_tiers[4][call_center_agent_uuid]" class="formfld" style="width: 200px">
			<option value=""></option>
		<option value="fd7b879d-a2dd-4ef7-9b05-06156a4f80b7">1601</option>
		<option value="6d66dadc-2d38-46e3-ba15-aa113ffa4892">1603</option>
		<option value="c2c0722e-0550-4caa-852e-b765791ef931">1604</option>
		<option value="6a256f0c-84a4-494a-8f66-2913e59779b6">1606</option>
		<option value="a342da12-0d53-4362-8238-417c7a50368c">1607</option>
		<option value="9557f6cf-785d-4e6e-ada6-b94910190556">1609</option>
		<option value="ef50191d-d238-4e44-a29c-db1c22634a21">1610</option>
		<option value="0e96e013-8199-47fd-8268-6c1dcd20aa96">1611</option>
		<option value="d1499462-3461-4d93-9441-72b3c24f857c">1612</option>
		<option value="78352cbe-c627-41cf-a212-a3fc8b8c6260">1613</option>
		<option value="305dbdf8-4546-4053-aa3e-fc268d5f9c91">1614</option>
		<option value="825a6bf5-809f-49fb-a197-7caffa87fd2d">1615</option>
		<option value="98d94ddc-8496-4b95-ae91-9ceee7380083">1616</option>
		<option value="b623066e-8587-4742-8727-a2228c3375d9">1617</option>
		<option value="b6a7e90d-2e94-4967-be9b-48d9c89db12e">1618</option>
		<option value="082eb3b0-fae5-4257-b538-be5e1f9156ca">1619</option>
		<option value="6d2ea436-325d-492c-9143-2b7298ad190a">1626</option>
		<option value="dfabcc60-538e-49dc-8deb-9ab5dedd8c36">1627</option>
		<option value="24d82d98-e94a-4c7f-b64a-feb05261ff6e">1628</option>
		<option value="a6862035-0524-48ee-921c-9b30dc31c2cf">1629</option>
		<option value="3961b4d6-210e-45e9-ac4f-a2ecb66913a5">1630</option>
		<option value="353d6cb2-dbec-44c1-995a-fe5d86e6e277">1631</option>
		<option value="30c7a03d-c175-431d-8c18-d74e1975cd0d">1632</option>
		<option value="fc9a2911-2a09-48ff-b89b-bda47b21998a">1634</option>
		<option value="6fe18ce4-8327-46a8-baab-0f228bcf5649">1635</option>
		<option value="6214f0c8-395c-4c59-87b5-002702195c51">1636</option>
		<option value="f8fba913-a39f-4ed2-94e5-6a51b931f80a">1637</option>
		<option value="7caa54c4-9899-46fb-8177-ebc1f83513d2">1638</option>
		<option value="2e92ca8c-9c72-40d3-9444-04da4bf50a6a">1639</option>
		<option value="f998c21b-f324-47c9-9624-a47876cf0a91">1641</option>
		<option value="d3aa8b8e-54e0-42e9-9f82-65e4fa996796">1642</option>
		<option value="e1f442e1-cbd9-4da9-a9db-2beaf61898b1">1644</option>
		<option value="52b282f8-6499-4b53-a92d-7407c2a601d0">1645</option>
		<option value="50f623e0-3092-428d-bc50-03f086867a82">1646</option>
		<option value="11940fa2-4bd9-433e-b4ad-5b8fb264422d">1647</option>
		<option value="d144d59d-da87-4640-8df5-ab44a4dee372">1649</option>
		<option value="13d3de44-422b-4cf7-aaca-dbb800dfd762">1650</option>
		<option value="864bec58-5173-40b4-8799-5995bb19cb2c">1651</option>
		<option value="6c504577-0c53-4aa5-8da1-c5ca00ebdcd6">1652</option>
		<option value="ac105f6c-6655-46b1-8adf-67e75b3fc6d5">1653</option>
		<option value="264d0e22-a4c8-42e9-947f-a67fab9f2fb9">1654</option>
		<option value="21ee25df-ad7e-4e82-83e7-379670220cdb">1655</option>
		<option value="5baa539e-ab37-4633-9351-42eff7c15edc">1656</option>
		<option value="91f77326-847d-4ed2-8bd1-92ee71180fd3">1657</option>
		<option value="2cdfd4fb-db22-4b31-93b0-febe5dfc61e3">1658</option>
		<option value="598ebf17-4471-475c-8451-7b25d3e1383e">1659</option>
		<option value="dabf7ba9-6d74-48df-b533-dda7219cea07">1660</option>
		<option value="05cf3919-cec9-4b12-9525-b492fa426da3">1661</option>
		<option value="92569c7d-f142-440f-a885-a86c72ad0936">1662</option>
		<option value="015cdc3e-26e7-4347-a2d7-483f0b88ec01">1663</option>
		<option value="1d20eec5-4f2f-43c8-bf2c-d06fb2e272f8">1664</option>
		<option value="97227560-247d-49b7-99c5-c19e3cac9552">1665</option>
		<option value="a958700d-5c31-48d7-9f87-e05cbebfc6c5">1666</option>
		<option value="1c063326-64e2-457c-876d-f536a8ba23a2">1668</option>
		<option value="19deda13-ef7f-43b7-a8c6-3d3f9601ca3c">1669</option>
		<option value="99b094b6-5e08-4472-adc9-853c604be499">1670</option>
		<option value="1265e427-d385-4a3e-a6dc-0093880cd1e5">1671</option>
		<option value="d796a830-cf23-4589-9b07-e610c6d2e7d8">1672</option>
		<option value="dbe0646d-7afb-4ba3-91b0-774952d1021b">1673</option>
		<option value="779905d7-fa83-412f-ad3e-433b53719835">1674</option>
		<option value="b3204faf-86f8-4530-97bd-a6aa22e4886c">1675</option>
		<option value="1bd17098-016f-44ab-b0d0-c6d42fbb944c">1676</option>
		<option value="f81bfa23-068e-41c1-b0b0-643652007a96">1677</option>
		<option value="898d2219-2288-40c6-b2bb-0ecf156cea95">1678</option>
		<option value="5b3d49a1-417d-46b6-b2f2-e6a2f97de078">1679</option>
		<option value="faca8f9b-4f10-403d-9e56-6ab0dae2acbd">1680</option>
		<option value="e9c17f2f-54e8-4a90-852e-fa7d5b1c4f1b">1681</option>
		<option value="9becac36-310b-4852-b464-062e88968c54">1682</option>
		<option value="ce2a697d-85e8-4bd7-a8b5-2f5b7906b883">1683</option>
		<option value="1dde0bf9-966a-40f4-8f02-68928048f0af">1684</option>
		<option value="77dadef4-03b9-4c8a-973e-2b912789a972">1685</option>
		<option value="d93b5abb-5dcf-4e40-af16-7338f5f64c3a">1686</option>
		<option value="92374956-dd15-4edb-a340-70139e0e941c">1687</option>
		<option value="e3923b66-6cf3-4da6-b1aa-caba339127c0">1688</option>
		<option value="567ae4fc-db87-4a4b-b405-265c50d55604">1689</option>
		<option value="7bd3f9e8-b321-4e92-a862-09d416ae1594">1690</option>
		<option value="38ad6832-537f-4f7b-af37-5b8f0f755359">1691</option>
		<option value="ac68ed3a-cf71-4d66-9577-3f6512fed3d1">1692</option>
		<option value="b7ad4660-4bd7-42a0-9223-f96bceac18b6">1693</option>
		<option value="9d3f1651-e68e-44b8-b425-2f697a0109af">1694</option>
		<option value="b70815c9-1672-414a-ae4c-25b39f54beff">1696</option>
		<option value="77cdd4d6-c15b-45c8-ad6c-671d8f9426b6">9998</option>
		<option value="b4ca43fc-d988-46d3-9cc6-2a8ecc5a0893">9999</option>
		</select>		</td>
<td class='' style='text-align: center;'>				 <select name="call_center_tiers[4][tier_level]" class="formfld">
		<option value="0" >0</option>
		<option value="1" >1</option>
		<option value="2" >2</option>
		<option value="3" >3</option>
		<option value="4" >4</option>
		<option value="5" >5</option>
		<option value="6" >6</option>
		<option value="7" >7</option>
		<option value="8" >8</option>
		<option value="9" >9</option>
		</select>
</td>
<td class='' style='text-align: center;'>
		<select name="call_center_tiers[4][tier_position]" class="formfld">
		<option value="0" >0</option>
		<option value="1" >1</option>
		<option value="2" >2</option>
		<option value="3" >3</option>
		<option value="4" >4</option>
		<option value="5" >5</option>
		<option value="6" >6</option>
		<option value="7" >7</option>
		<option value="8" >8</option>
		<option value="9" >9</option>
		</select>
</td>
<td class=''>			<a href="call_center_queue_edit.php?id=988348b3-465d-449c-8137-ec3031941c84&call_center_tier_uuid=d3f3f74b-328c-485e-a751-9f6999567c95&a=delete" alt="Delete" onclick="return confirm('Do you really want to DELETE this?');"><button type='button' class='btn btn-default list_control_icon'><span class='fas fa-minus'></span></button></a>		</td>
</tr>
</table>
<br>
Tiers assign agents to queues.
<br />
</td></tr><tr>
<td class='vncell' valign='top' align='left' nowrap>
Music on Hold
</td>
<td class='vtable' align='left'>
<select class='formfld' name='queue_moh_sound' id='queue_moh_sound' style='width: auto;'>
<option value=''></option>
<optgroup label='Music on Hold'>
<option value='local_stream://default' >default</option>
</optgroup>
<optgroup label='Ringtones'>		<option value='${au-ring}'>au-ring</option>
<option value='${be-ring}'>be-ring</option>
<option value='${bong-ring}'>bong-ring</option>
<option value='${ca-ring}'>ca-ring</option>
<option value='${cn-ring}'>cn-ring</option>
<option value='${cy-ring}'>cy-ring</option>
<option value='${cz-ring}'>cz-ring</option>
<option value='${de-ring}'>de-ring</option>
<option value='${dk-ring}'>dk-ring</option>
<option value='${dz-ring}'>dz-ring</option>
<option value='${eg-ring}'>eg-ring</option>
<option value='${fi-ring}'>fi-ring</option>
<option value='${fr-ring}'>fr-ring</option>
<option value='${hk-ring}'>hk-ring</option>
<option value='${hu-ring}'>hu-ring</option>
<option value='${il-ring}'>il-ring</option>
<option value='${in-ring}'>in-ring</option>
<option value='${it-ring}'>it-ring</option>
<option value='${jp-ring}'>jp-ring</option>
<option value='${ko-ring}'>ko-ring</option>
<option value='${pk-ring}'>pk-ring</option>
<option value='${pl-ring}'>pl-ring</option>
<option value='${pt-ring}' selected="selected">pt-ring</option>
<option value='${ro-ring}'>ro-ring</option>
<option value='${rs-ring}'>rs-ring</option>
<option value='${ru-ring}'>ru-ring</option>
<option value='${sa-ring}'>sa-ring</option>
<option value='${tr-ring}'>tr-ring</option>
<option value='${uk-ring}'>uk-ring</option>
<option value='${us-ring}'>us-ring</option>
<option value='silence'>Silence</option>
</optgroup>
<optgroup label='Tones'>		<option value='${bong-us-tone}'>bong-us-tone</option>
<option value='${busy-au-tone}'>busy-au-tone</option>
<option value='${busy-us-tone}'>busy-us-tone</option>
<option value='${vacant-uk-tone}'>vacant-uk-tone</option>
<option value='${vacant-us-tone}'>vacant-us-tone</option>
</optgroup>
</select>
<br />
Select the desired hold music.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Record
</td>
<td class='vtable' align='left'>
<select class='formfld' name='queue_record_template'>
<option value='&sol;var&sol;lib&sol;freeswitch&sol;recordings&sol;voip01&period;dial&period;opssbank&period;com&period;br&sol;archive&sol;&dollar;&lbrace;strftime&lpar;&percnt;Y&rpar;&rcub;&sol;&dollar;&lbrace;strftime&lpar;&percnt;b&rpar;&rcub;&sol;&dollar;&lbrace;strftime&lpar;&percnt;d&rpar;&rcub;&sol;&dollar;&lbrace;uuid&rcub;&period;&dollar;&lbrace;record&lowbar;ext&rcub;' selected='selected' >True</option>
<option value=''>False</option>
</select>
<br />
Save the recording.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Time Base Score
</td>
<td class='vtable' align='left'>
<select class='formfld' name='queue_time_base_score'>
<option value='system' selected='selected' >System</option>
<option value='queue'>Queue</option>
</select>
<br />
Select the time base score.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Time Base Score Seconds
</td>
<td class='vtable' align='left'>
<input class='formfld' type='number' name='queue_time_base_score_sec' maxlength='255' min='0' step='1' value=''>
<br />
Set the time base score in seconds. Higher numbers mean higher priority.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Max Wait Time
</td>
<td class='vtable' align='left'>
<input class='formfld' type='number' name='queue_max_wait_time' maxlength='255' min='0' step='1' value='0'>
<br />
Enter the max wait time.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Max Wait Time with No Agent
</td>
<td class='vtable' align='left'>
<input class='formfld' type='number' name='queue_max_wait_time_with_no_agent' maxlength='255' min='0' step='1' value='90'>
<br />
Enter the max wait time with no agent.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Max Wait Time with No Agent Time Reached
</td>
<td class='vtable' align='left'>
<input class='formfld' type='number' name='queue_max_wait_time_with_no_agent_time_reached' maxlength='255' min='0' step='1' value='30'>
<br />
Enter the max wait time with no agent time reached.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Timeout Action
</td>
<td class='vtable' align='left'>

<script>
var Objs;

function changeToInputqueue_timeout_action(obj){
tb=document.createElement('INPUT');
tb.type='text';
tb.name=obj.name;
tb.className='formfld';
tb.setAttribute('id', 'queue_timeout_action');
tb.setAttribute('style', 'width: 200px;');
tb.value=obj.options[obj.selectedIndex].value;
document.getElementById('btn_select_to_input_queue_timeout_action').style.visibility = 'hidden';
tbb=document.createElement('INPUT');
tbb.setAttribute('class', 'btn');
tbb.setAttribute('style', 'margin-left: 4px;');
tbb.type='button';
tbb.value=$('<div />').html('&#9665;').text();
tbb.objs=[obj,tb,tbb];
tbb.onclick=function(){ Replacequeue_timeout_action(this.objs); }
obj.parentNode.insertBefore(tb,obj);
obj.parentNode.insertBefore(tbb,obj);
obj.parentNode.removeChild(obj);
Replacequeue_timeout_action(this.objs);
}

function Replacequeue_timeout_action(obj){
obj[2].parentNode.insertBefore(obj[0],obj[2]);
obj[0].parentNode.removeChild(obj[1]);
obj[0].parentNode.removeChild(obj[2]);
document.getElementById('btn_select_to_input_queue_timeout_action').style.visibility = 'visible';
}
</script>

<select name='queue_timeout_action' id='queue_timeout_action' class='formfld' style='width: 200px;' onchange="">
	<option value=''></option>
<optgroup label='Call Center'>
	<option value='transfer&colon;55550 XML voip01&period;dial&period;opssbank&period;com&period;br' >55550 ATENDIMENTO</option>
	<option value='transfer&colon;808080 XML voip01&period;dial&period;opssbank&period;com&period;br' >808080 AT&lowbar;CRM</option>
	<option value='transfer&colon;200110 XML voip01&period;dial&period;opssbank&period;com&period;br' >200110 Admin Opss</option>
	<option value='transfer&colon;200111 XML voip01&period;dial&period;opssbank&period;com&period;br' >200111 CAMPANHAS DE MKT</option>
	<option value='transfer&colon;66666 XML voip01&period;dial&period;opssbank&period;com&period;br' >66666 CRM</option>
	<option value='transfer&colon;200100 XML voip01&period;dial&period;opssbank&period;com&period;br' >200100 Cavalcante &amp; Mendes</option>
	<option value='transfer&colon;200101 XML voip01&period;dial&period;opssbank&period;com&period;br' >200101 Cedro 7 Seguros</option>
	<option value='transfer&colon;200102 XML voip01&period;dial&period;opssbank&period;com&period;br' >200102 Davi Miret</option>
	<option value='transfer&colon;200104 XML voip01&period;dial&period;opssbank&period;com&period;br' >200104 FIRMASEG</option>
	<option value='transfer&colon;200105 XML voip01&period;dial&period;opssbank&period;com&period;br' >200105 FP 2</option>
	<option value='transfer&colon;200103 XML voip01&period;dial&period;opssbank&period;com&period;br' >200103 Felice Adm</option>
	<option value='transfer&colon;200106 XML voip01&period;dial&period;opssbank&period;com&period;br' >200106 GHGM 1</option>
	<option value='transfer&colon;200112 XML voip01&period;dial&period;opssbank&period;com&period;br' >200112 Home Office</option>
	<option value='transfer&colon;200107 XML voip01&period;dial&period;opssbank&period;com&period;br' >200107 Inonvax</option>
	<option value='transfer&colon;200108 XML voip01&period;dial&period;opssbank&period;com&period;br' >200108 Ivaro Matos</option>
	<option value='transfer&colon;200109 XML voip01&period;dial&period;opssbank&period;com&period;br' >200109 Nascional 1</option>
	<option value='transfer&colon;200113 XML voip01&period;dial&period;opssbank&period;com&period;br' >200113 Opss Lab</option>
	<option value='transfer&colon;200114 XML voip01&period;dial&period;opssbank&period;com&period;br' >200114 Opss Team</option>
	<option value='transfer&colon;200116 XML voip01&period;dial&period;opssbank&period;com&period;br' >200116 Prevseguros</option>
	<option value='transfer&colon;200118 XML voip01&period;dial&period;opssbank&period;com&period;br' >200118 RF Feijo</option>
	<option value='transfer&colon;200117 XML voip01&period;dial&period;opssbank&period;com&period;br' >200117 RIB COR</option>
	<option value='transfer&colon;200119 XML voip01&period;dial&period;opssbank&period;com&period;br' >200119 Ronaldo Alex Seguros</option>
	<option value='transfer&colon;200115 XML voip01&period;dial&period;opssbank&period;com&period;br' >200115 Teste p&eacute;ricles</option>
	<option value='transfer&colon;200220 XML voip01&period;dial&period;opssbank&period;com&period;br' >200220 Victor Hugo Ferreira</option>
	<option value='transfer&colon;200221 XML voip01&period;dial&period;opssbank&period;com&period;br' >200221 Wilson Cor</option>
	<option value='transfer&colon;200170 XML voip01&period;dial&period;opssbank&period;com&period;br' >200170 c&lowbar;170</option>
</optgroup>
<optgroup label='Call Groups'>
	<option value='bridge&colon;group&sol;sales&commat;voip01&period;dial&period;opssbank&period;com&period;br' >sales</option>
	<option value='bridge&colon;group&sol;support&commat;voip01&period;dial&period;opssbank&period;com&period;br' >support</option>
</optgroup>
<optgroup label='Extensions'>
	<option value='transfer&colon;1467 XML voip01&period;dial&period;opssbank&period;com&period;br' >1467 Rogerio Pessoa</option>
	<option value='transfer&colon;1601 XML voip01&period;dial&period;opssbank&period;com&period;br' >1601 Danilo Machado</option>
	<option value='transfer&colon;1602 XML voip01&period;dial&period;opssbank&period;com&period;br' >1602 1602WEBRTC</option>
	<option value='transfer&colon;1603 XML voip01&period;dial&period;opssbank&period;com&period;br' >1603 1603descricao</option>
	<option value='transfer&colon;1604 XML voip01&period;dial&period;opssbank&period;com&period;br' >1604 Marcelo Palmieri</option>
	<option value='transfer&colon;1606 XML voip01&period;dial&period;opssbank&period;com&period;br' >1606 Gustavo Felipe</option>
	<option value='transfer&colon;1607 XML voip01&period;dial&period;opssbank&period;com&period;br' >1607 Franciane Kelly</option>
	<option value='transfer&colon;1608 XML voip01&period;dial&period;opssbank&period;com&period;br' >1608 Leonardo Mendes</option>
	<option value='transfer&colon;1609 XML voip01&period;dial&period;opssbank&period;com&period;br' >1609 Wesley Jose</option>
	<option value='transfer&colon;1610 XML voip01&period;dial&period;opssbank&period;com&period;br' >1610 Marcos Machado</option>
	<option value='transfer&colon;1611 XML voip01&period;dial&period;opssbank&period;com&period;br' >1611 Vitor Junior</option>
	<option value='transfer&colon;1612 XML voip01&period;dial&period;opssbank&period;com&period;br' >1612 Alexandre Dias</option>
	<option value='transfer&colon;1613 XML teste&lowbar;algar' >1613 1613</option>
	<option value='transfer&colon;1614 XML voip01&period;dial&period;opssbank&period;com&period;br' >1614 Gabriela Araujo</option>
	<option value='transfer&colon;1615 XML voip01&period;dial&period;opssbank&period;com&period;br' >1615 Michell Augusto</option>
	<option value='transfer&colon;1616 XML voip01&period;dial&period;opssbank&period;com&period;br' >1616 Mateus Assis</option>
	<option value='transfer&colon;1617 XML voip01&period;dial&period;opssbank&period;com&period;br' >1617 1617</option>
	<option value='transfer&colon;1618 XML voip01&period;dial&period;opssbank&period;com&period;br' >1618 Andre Ventura</option>
	<option value='transfer&colon;1619 XML voip01&period;dial&period;opssbank&period;com&period;br' >1619 1619</option>
	<option value='transfer&colon;1620 XML voip01&period;dial&period;opssbank&period;com&period;br' >1620 POSITIVA ESTHER MARTINS</option>
	<option value='transfer&colon;1621 XML voip01&period;dial&period;opssbank&period;com&period;br' >1621 POSITIVA GABRIELA SANTOS</option>
	<option value='transfer&colon;1622 XML voip01&period;dial&period;opssbank&period;com&period;br' >1622 POSITIVA JULIA BEATRIZ</option>
	<option value='transfer&colon;1623 XML voip01&period;dial&period;opssbank&period;com&period;br' >1623 POSITIVA MARIA GABRIELLI</option>
	<option value='transfer&colon;1624 XML voip01&period;dial&period;opssbank&period;com&period;br' >1624 POSITIVA IEDA CRISTINA</option>
	<option value='transfer&colon;1625 XML voip01&period;dial&period;opssbank&period;com&period;br' >1625 Positiva Victoria</option>
	<option value='transfer&colon;1626 XML voip01&period;dial&period;opssbank&period;com&period;br' >1626 Biane Tatielly</option>
	<option value='transfer&colon;1627 XML voip01&period;dial&period;opssbank&period;com&period;br' >1627 Marcos Machado</option>
	<option value='transfer&colon;1628 XML voip01&period;dial&period;opssbank&period;com&period;br' >1628 Regiane Oliveira</option>
	<option value='transfer&colon;1629 XML voip01&period;dial&period;opssbank&period;com&period;br' >1629 Stephanie Ramos</option>
	<option value='transfer&colon;1630 XML voip01&period;dial&period;opssbank&period;com&period;br' >1630 Vitor Junio</option>
	<option value='transfer&colon;1631 XML voip01&period;dial&period;opssbank&period;com&period;br' >1631 Viviane Adame</option>
	<option value='transfer&colon;1632 XML voip01&period;dial&period;opssbank&period;com&period;br' >1632 Diogo Machado</option>
	<option value='transfer&colon;1633 XML voip01&period;dial&period;opssbank&period;com&period;br' >1633 Rodrigo Mattos</option>
	<option value='transfer&colon;1634 XML voip01&period;dial&period;opssbank&period;com&period;br' >1634 Augusto Nascimento</option>
	<option value='transfer&colon;1635 XML voip01&period;dial&period;opssbank&period;com&period;br' >1635 Douglas Machado</option>
	<option value='transfer&colon;1636 XML voip01&period;dial&period;opssbank&period;com&period;br' >1636 Eduarda Guedes</option>
	<option value='transfer&colon;1637 XML voip01&period;dial&period;opssbank&period;com&period;br' >1637 1637</option>
	<option value='transfer&colon;1641 XML voip01&period;dial&period;opssbank&period;com&period;br' >1641 Debora Menezes</option>
	<option value='transfer&colon;1642 XML voip01&period;dial&period;opssbank&period;com&period;br' >1642 Inovax David Jesus</option>
	<option value='transfer&colon;1643 XML voip01&period;dial&period;opssbank&period;com&period;br' >1643 Inovax Julia de Souza</option>
	<option value='transfer&colon;1644 XML voip01&period;dial&period;opssbank&period;com&period;br' >1644 Ana Julia</option>
	<option value='transfer&colon;1645 XML voip01&period;dial&period;opssbank&period;com&period;br' >1645 Lorraine Barbosa</option>
	<option value='transfer&colon;1646 XML voip01&period;dial&period;opssbank&period;com&period;br' >1646 Natalia Duarte</option>
	<option value='transfer&colon;1647 XML voip01&period;dial&period;opssbank&period;com&period;br' >1647 Ronaldo Alex</option>
	<option value='transfer&colon;1649 XML voip01&period;dial&period;opssbank&period;com&period;br' >1649 1649</option>
	<option value='transfer&colon;1650 XML voip01&period;dial&period;opssbank&period;com&period;br' >1650 1650</option>
	<option value='transfer&colon;1651 XML voip01&period;dial&period;opssbank&period;com&period;br' >1651 1651</option>
	<option value='transfer&colon;1652 XML voip01&period;dial&period;opssbank&period;com&period;br' >1652 1652</option>
	<option value='transfer&colon;1653 XML voip01&period;dial&period;opssbank&period;com&period;br' >1653 1653</option>
	<option value='transfer&colon;1654 XML voip01&period;dial&period;opssbank&period;com&period;br' >1654 1654</option>
	<option value='transfer&colon;1655 XML voip01&period;dial&period;opssbank&period;com&period;br' >1655 1655</option>
	<option value='transfer&colon;1656 XML voip01&period;dial&period;opssbank&period;com&period;br' >1656 Gleiciane Rodrigues</option>
	<option value='transfer&colon;1657 XML voip01&period;dial&period;opssbank&period;com&period;br' >1657 Vivian Fernandes</option>
	<option value='transfer&colon;1658 XML voip01&period;dial&period;opssbank&period;com&period;br' >1658 Erilaine&nbsp;Santos</option>
	<option value='transfer&colon;1659 XML voip01&period;dial&period;opssbank&period;com&period;br' >1659 Fladielle Santos</option>
	<option value='transfer&colon;1660 XML voip01&period;dial&period;opssbank&period;com&period;br' >1660 Ana Neres</option>
	<option value='transfer&colon;1661 XML voip01&period;dial&period;opssbank&period;com&period;br' >1661 1661</option>
	<option value='transfer&colon;1662 XML voip01&period;dial&period;opssbank&period;com&period;br' >1662 1662</option>
	<option value='transfer&colon;1663 XML voip01&period;dial&period;opssbank&period;com&period;br' >1663 1663</option>
	<option value='transfer&colon;1664 XML voip01&period;dial&period;opssbank&period;com&period;br' >1664 1664</option>
	<option value='transfer&colon;1665 XML voip01&period;dial&period;opssbank&period;com&period;br' >1665 1665</option>
	<option value='transfer&colon;1666 XML voip01&period;dial&period;opssbank&period;com&period;br' >1666 1666</option>
	<option value='transfer&colon;1667 XML voip01&period;dial&period;opssbank&period;com&period;br' >1667 Joao Victor</option>
	<option value='transfer&colon;1668 XML voip01&period;dial&period;opssbank&period;com&period;br' >1668 Maria Eduarda</option>
	<option value='transfer&colon;1669 XML voip01&period;dial&period;opssbank&period;com&period;br' >1669 Bruna Antonia</option>
	<option value='transfer&colon;1670 XML voip01&period;dial&period;opssbank&period;com&period;br' >1670 Daniel Kesley</option>
	<option value='transfer&colon;1671 XML voip01&period;dial&period;opssbank&period;com&period;br' >1671 Cecilia Aparecida</option>
	<option value='transfer&colon;1672 XML voip01&period;dial&period;opssbank&period;com&period;br' >1672 Elisangela Alves</option>
	<option value='transfer&colon;1673 XML voip01&period;dial&period;opssbank&period;com&period;br' >1673 Nadia Priscila</option>
	<option value='transfer&colon;1674 XML voip01&period;dial&period;opssbank&period;com&period;br' >1674 Analice Soares</option>
	<option value='transfer&colon;1676 XML voip01&period;dial&period;opssbank&period;com&period;br' >1676 1676</option>
	<option value='transfer&colon;1677 XML voip01&period;dial&period;opssbank&period;com&period;br' >1677 1677</option>
	<option value='transfer&colon;1678 XML voip01&period;dial&period;opssbank&period;com&period;br' >1678 1678</option>
	<option value='transfer&colon;1679 XML voip01&period;dial&period;opssbank&period;com&period;br' >1679 1679</option>
	<option value='transfer&colon;1680 XML voip01&period;dial&period;opssbank&period;com&period;br' >1680 1680</option>
	<option value='transfer&colon;1681 XML voip01&period;dial&period;opssbank&period;com&period;br' >1681 1681</option>
	<option value='transfer&colon;1682 XML voip01&period;dial&period;opssbank&period;com&period;br' >1682 1682</option>
	<option value='transfer&colon;1683 XML voip01&period;dial&period;opssbank&period;com&period;br' >1683 1683</option>
	<option value='transfer&colon;1684 XML voip01&period;dial&period;opssbank&period;com&period;br' >1684 1684</option>
	<option value='transfer&colon;1685 XML voip01&period;dial&period;opssbank&period;com&period;br' >1685 1685</option>
	<option value='transfer&colon;1686 XML voip01&period;dial&period;opssbank&period;com&period;br' >1686 1686</option>
	<option value='transfer&colon;1687 XML voip01&period;dial&period;opssbank&period;com&period;br' >1687 1687</option>
	<option value='transfer&colon;1688 XML voip01&period;dial&period;opssbank&period;com&period;br' >1688 1688</option>
	<option value='transfer&colon;1689 XML voip01&period;dial&period;opssbank&period;com&period;br' >1689 1689</option>
	<option value='transfer&colon;1690 XML voip01&period;dial&period;opssbank&period;com&period;br' >1690 1690</option>
	<option value='transfer&colon;1691 XML voip01&period;dial&period;opssbank&period;com&period;br' >1691 1691</option>
	<option value='transfer&colon;1692 XML voip01&period;dial&period;opssbank&period;com&period;br' >1692 1692</option>
	<option value='transfer&colon;1693 XML voip01&period;dial&period;opssbank&period;com&period;br' >1693 1693</option>
	<option value='transfer&colon;1694 XML voip01&period;dial&period;opssbank&period;com&period;br' >1694 1694</option>
	<option value='transfer&colon;1695 XML voip01&period;dial&period;opssbank&period;com&period;br' >1695 1695</option>
	<option value='transfer&colon;1696 XML voip01&period;dial&period;opssbank&period;com&period;br' >1696 1696</option>
	<option value='transfer&colon;1697 XML voip01&period;dial&period;opssbank&period;com&period;br' >1697</option>
	<option value='transfer&colon;7777 XML escritorio' >7777 CHARLESTON</option>
	<option value='transfer&colon;8887 XML escritorio' >8887 RAMAL CALL CENTER</option>
	<option value='transfer&colon;8888 XML escritorio' >8888 RAMAL ESCRITORIO</option>
	<option value='transfer&colon;9996 XML discador&lowbar;externo' >9996</option>
	<option value='transfer&colon;9997 XML voip01&period;dial&period;opssbank&period;com&period;br' >9997</option>
	<option value='transfer&colon;9998 XML voip01&period;dial&period;opssbank&period;com&period;br' >9998</option>
	<option value='transfer&colon;9999 XML voip01&period;dial&period;opssbank&period;com&period;br' >9999</option>
</optgroup>
<optgroup label='Tones'>
	<option value='playback&colon;tone&lowbar;stream&colon;&sol;&sol;v&equals;-7&semi;&percnt;&lpar;100&comma;0&comma;941&period;0&comma;1477&period;0&rpar;&semi;v&equals;-7&semi;&gt;&equals;2&semi;&plus;&equals;&period;1&semi;&percnt;&lpar;1400&comma;0&comma;350&comma;440&rpar;' >bong-us-tone</option>
	<option value='playback&colon;tone&lowbar;stream&colon;&sol;&sol;v&equals;-13&semi;&percnt;&lpar;375&comma;375&comma;420&rpar;&semi;v&equals;-23&semi;&percnt;&lpar;375&comma;375&comma;420&rpar;' >busy-au-tone</option>
	<option value='playback&colon;tone&lowbar;stream&colon;&sol;&sol;&percnt;&lpar;500&comma;500&comma;480&comma;620&rpar;' >busy-us-tone</option>
	<option value='playback&colon;tone&lowbar;stream&colon;&sol;&sol;&percnt;&lpar;330&comma;15&comma;950&rpar;&semi;&percnt;&lpar;330&comma;15&comma;1400&rpar;&semi;&percnt;&lpar;330&comma;1000&comma;1800&rpar;' >vacant-uk-tone</option>
	<option value='playback&colon;tone&lowbar;stream&colon;&sol;&sol;&percnt;&lpar;274&comma;0&comma;913&period;8&rpar;&semi;&percnt;&lpar;274&comma;0&comma;1370&period;6&rpar;&semi;&percnt;&lpar;380&comma;0&comma;1776&period;7&rpar;' >vacant-us-tone</option>
</optgroup>
<optgroup label='Voicemails'>
	<option value='transfer&colon;&ast;991608 XML voip01&period;dial&period;opssbank&period;com&period;br' >1608 Leonardo Mendes &#9993</option>
	<option value='transfer&colon;&ast;991613 XML voip01&period;dial&period;opssbank&period;com&period;br' >1613 1613 &#9993</option>
	<option value='transfer&colon;&ast;991638 XML voip01&period;dial&period;opssbank&period;com&period;br' >1638 Camila Moreira &#9993</option>
	<option value='transfer&colon;&ast;991639 XML voip01&period;dial&period;opssbank&period;com&period;br' >1639 Gabrielle Moreira &#9993</option>
	<option value='transfer&colon;&ast;991640 XML voip01&period;dial&period;opssbank&period;com&period;br' >1640 Ian &#9993</option>
	<option value='transfer&colon;&ast;991641 XML voip01&period;dial&period;opssbank&period;com&period;br' >1641 Debora Menezes &#9993</option>
	<option value='transfer&colon;&ast;991644 XML voip01&period;dial&period;opssbank&period;com&period;br' >1644 Ana Julia &#9993</option>
	<option value='transfer&colon;&ast;991645 XML voip01&period;dial&period;opssbank&period;com&period;br' >1645 Lorraine Barbosa &#9993</option>
	<option value='transfer&colon;&ast;991646 XML voip01&period;dial&period;opssbank&period;com&period;br' >1646 Natalia Duarte &#9993</option>
	<option value='transfer&colon;&ast;991647 XML voip01&period;dial&period;opssbank&period;com&period;br' >1647 Ronaldo Alex &#9993</option>
	<option value='transfer&colon;&ast;991648 XML voip01&period;dial&period;opssbank&period;com&period;br' >1648 Miriam Oliveira &#9993</option>
	<option value='transfer&colon;&ast;991656 XML voip01&period;dial&period;opssbank&period;com&period;br' >1656 Gleiciane Rodrigues &#9993</option>
	<option value='transfer&colon;&ast;991657 XML voip01&period;dial&period;opssbank&period;com&period;br' >1657 Vivian Fernandes &#9993</option>
	<option value='transfer&colon;&ast;991658 XML voip01&period;dial&period;opssbank&period;com&period;br' >1658 Erilaine&nbsp;Santos &#9993</option>
	<option value='transfer&colon;&ast;991659 XML voip01&period;dial&period;opssbank&period;com&period;br' >1659 Fladielle Santos &#9993</option>
	<option value='transfer&colon;&ast;991660 XML voip01&period;dial&period;opssbank&period;com&period;br' >1660 Ana Neres &#9993</option>
	<option value='transfer&colon;&ast;991667 XML voip01&period;dial&period;opssbank&period;com&period;br' >1667 Joao Victor &#9993</option>
	<option value='transfer&colon;&ast;991668 XML voip01&period;dial&period;opssbank&period;com&period;br' >1668 Maria Eduarda &#9993</option>
	<option value='transfer&colon;&ast;991669 XML voip01&period;dial&period;opssbank&period;com&period;br' >1669 Bruna Antonia &#9993</option>
	<option value='transfer&colon;&ast;991670 XML voip01&period;dial&period;opssbank&period;com&period;br' >1670 Daniel Kesley &#9993</option>
	<option value='transfer&colon;&ast;991671 XML voip01&period;dial&period;opssbank&period;com&period;br' >1671 Cecilia Aparecida &#9993</option>
	<option value='transfer&colon;&ast;991672 XML voip01&period;dial&period;opssbank&period;com&period;br' >1672 Elisangela Alves &#9993</option>
	<option value='transfer&colon;&ast;991673 XML voip01&period;dial&period;opssbank&period;com&period;br' >1673 Nadia Priscila &#9993</option>
	<option value='transfer&colon;&ast;991674 XML voip01&period;dial&period;opssbank&period;com&period;br' >1674 Analice Soares &#9993</option>
	<option value='transfer&colon;&ast;991675 XML voip01&period;dial&period;opssbank&period;com&period;br' >1675 Silvania Almeida &#9993</option>
	<option value='transfer&colon;&ast;999996 XML voip01&period;dial&period;opssbank&period;com&period;br' >9996  &#9993</option>
	<option value='transfer&colon;&ast;999997 XML voip01&period;dial&period;opssbank&period;com&period;br' >9997  &#9993</option>
	<option value='transfer&colon;&ast;999998 XML voip01&period;dial&period;opssbank&period;com&period;br' >9998  &#9993</option>
	<option value='transfer&colon;&ast;999999 XML voip01&period;dial&period;opssbank&period;com&period;br' >9999  &#9993</option>
</optgroup>
<optgroup label='Other'>
	<option value='transfer&colon;&ast;98 XML voip01&period;dial&period;opssbank&period;com&period;br' >Check Voicemail</option>
	<option value='transfer&colon;&ast;411 XML voip01&period;dial&period;opssbank&period;com&period;br' >Company Directory</option>
	<option value='hangup&colon;' >Hangup</option>
	<option value='transfer&colon;&ast;732 XML voip01&period;dial&period;opssbank&period;com&period;br' >Record</option>
</optgroup>
</select>
<input type='button' id='btn_select_to_input_queue_timeout_action' class='btn' name='' alt='back' onclick='changeToInputqueue_timeout_action(document.getElementById("queue_timeout_action"));this.style.visibility = "hidden";' value='&#9665;'><br />
Set the action to perform when the max wait time is reached.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Tier Rules Apply
</td>
<td class='vtable' align='left'>
<select class='formfld' name='queue_tier_rules_apply'>
<option value='true'>True</option>
<option value='false' selected='selected' >False</option>
</select>
<br />
Set the tier rule rules apply to true or false.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Tier Rule Wait Second
</td>
<td class='vtable' align='left'>
<input class='formfld' type='number' name='queue_tier_rule_wait_second' maxlength='255' min='0' step='1' value='30'>
<br />
Enter the tier rule wait seconds.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Tier Rule Wait Multiply Level
</td>
<td class='vtable' align='left'>
<select class='formfld' name='queue_tier_rule_wait_multiply_level'>
<option value='true' selected='selected' >True</option>
<option value='false'>False</option>
</select>
<br />
Set the tier rule wait multiply level to true or false.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Tier Rule No Agent No Wait
</td>
<td class='vtable' align='left'>
<select class='formfld' name='queue_tier_rule_no_agent_no_wait'>
<option value='true' selected='selected' >True</option>
<option value='false'>False</option>
</select>
<br />
Enter the tier rule no agent no wait.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Discard Abandoned After
</td>
<td class='vtable' align='left'>
<input class='formfld' type='number' name='queue_discard_abandoned_after' maxlength='255' min='0' step='1' value='900'>
<br />
The number of seconds before the abandoned call is removed from the queue.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Abandoned Resume Allowed
</td>
<td class='vtable' align='left'>
<select class='formfld' name='queue_abandoned_resume_allowed'>
<option value='false' selected='selected' >False</option>
<option value='true'>True</option>
</select>
<br />
A caller who has left the queue can resume their position in the queue by calling again before the abandoned call has been discarded.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Caller ID Name Prefix
</td>
<td class='vtable' align='left'>
<input class='formfld' type='text' name='queue_cid_prefix' maxlength='255' value=''>
<br />
Set a prefix on the caller ID name.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Announce Sound
</td>
<td class='vtable' align='left'>
<script>
var objs;

function changeToInputqueue_announce_sound(obj){
tb=document.createElement('INPUT');
tb.type='text';
tb.name=obj.name;
tb.className='formfld';
tb.setAttribute('id', 'queue_announce_sound');
tb.setAttribute('style', '');
tb.value=obj.options[obj.selectedIndex].value;
document.getElementById('btn_select_to_input_queue_announce_sound').style.visibility = 'hidden';
tbb=document.createElement('INPUT');
tbb.setAttribute('class', 'btn');
tbb.setAttribute('style', 'margin-left: 4px;');
tbb.type='button';
tbb.value=$('<div />').html('&#9665;').text();
tbb.objs=[obj,tb,tbb];
tbb.onclick=function(){ Replacequeue_announce_sound(this.objs); }
obj.parentNode.insertBefore(tb,obj);
obj.parentNode.insertBefore(tbb,obj);
obj.parentNode.removeChild(obj);
Replacequeue_announce_sound(this.objs);
}

function Replacequeue_announce_sound(obj){
obj[2].parentNode.insertBefore(obj[0],obj[2]);
obj[0].parentNode.removeChild(obj[1]);
obj[0].parentNode.removeChild(obj[2]);
document.getElementById('btn_select_to_input_queue_announce_sound').style.visibility = 'visible';
}
</script>

<select name='queue_announce_sound' id='queue_announce_sound' class='formfld'>
<option></option>
</select>
<input type='button' id='btn_select_to_input_queue&lowbar;announce&lowbar;sound' class='btn' name='' alt='back' onclick='changeToInputqueue&lowbar;announce&lowbar;sound(document.getElementById("queue&lowbar;announce&lowbar;sound"));this.style.visibility = "hidden";' value='&#9665;'>	<br />
A sound to play for a caller at specific intervals, as defined in seconds by the Announce Frequency. Full path to the recording is required.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Announce Frequency
</td>
<td class='vtable' align='left'>
<input class='formfld' type='number' name='queue_announce_frequency' maxlength='255' min='0' step='1' value=''>
<br />
How often should we play the announce sound. Enter a number in seconds
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Exit Key
</td>
<td class='vtable' align='left'>
<input class='formfld' type='text' name='queue_cc_exit_keys' value=''>
<br />
Define a key that can be used to exit the queue.
</td>
</tr>
<tr>
<td class='vncell' valign='top' align='left' nowrap>
Description
</td>
<td class='vtable' align='left'>
<input class='formfld' type='text' name='queue_description' maxlength='255' value="">
<br />
Enter the description.
</td>
</tr>
</table><br><br><input type='hidden' name='call_center_queue_uuid' value='988348b3-465d-449c-8137-ec3031941c84'>
<input type='hidden' name='dialplan_uuid' value='dfd61123-d156-42da-a8f7-a448b0f39a8a'>
<input type='hidden' name='872eebdfceb7f762f5b8aebeceea7b177cc029aa1154d5b0a795d1eb65d7c690' value='80d9e6f11faccc3c90c6d1a9e473721598ea57455ba11c232b7c86ca7945ea17'>
</form>
	</div>
	<div id='footer'>
		<span class='footer'>&copy; Copyright 2008 - 2024 <a href='http://www.fusionpbx.com' class='footer' target='_blank'>fusionpbx.com</a> All rights reserved.</span>
	</div>
	</div>

</body>
</html>
`
}
