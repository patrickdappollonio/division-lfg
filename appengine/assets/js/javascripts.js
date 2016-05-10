(function() {
	"use strict";

	// Vuejs configuration & Vue-resource configuration
	Vue.config.devtools = true;
	Vue.config.delimiters = ['@{', '}'];
	Vue.http.headers.common['X-CSRF-Token'] = document.querySelector('meta[name="csrf-token"]').getAttribute('content');

	// The URLs for each endpoint
	var urls = {
		"search": document.querySelector('#searchform').getAttribute('data-action'),
		"addnew": document.querySelector('#add-new-agent').getAttribute('data-action')
	};

	// Regular expressions for verification
	var regexs = {
		"username": /([\w \-\_\.]){3,20}/,
	};

	// Pretty-print the stat values
	Vue.filter('hundreds', function(value) {
		if (value >= 1000) {
			return (value / 1000).toFixed(0) + "K";
		} else {
			return value;
		}
	});

	// Register the agentblock component
	Vue.component('agent-block', {
		props: ['agent'],
		template: '#agent-block-template'
	});

	// Create the Vuejs app instance
	var app = new Vue({
		el: '#app',

		methods: {
			// Search all available agents using the sort-of API we have, plus
			// includes a bit of security by using CSRF tokens
			searchAgents: function() {
				// Send the search request to the endpoint
				this.$http.post(urls.search, this.searchparams).then(function(response) {
					if (!!response.data && !!response.data.agents) {
						this.$set('agents', response.data.agents);
					}
				},
				function(err) {
					console.log(err);
				});
			},

			// Function to add new agents to the system
			addNewAgent: function() {
				// Check if the form is valid
				if (!this.isAgentFormValid) {
					this.$set('shouldShowAgentValidation', true);
					$('#modal-add-agent').animate({ scrollTop: 0 }, 'fast');
					return;
				}

				// Send the data to the API endpoint
				this.$http.post(urls.addnew, this.agentinfo).then(function(response) {
					console.log(response);

					// TODO
					// 1. Clean the this.agentinfo with default values
					// 2. The API will return an agent object, append it to this.agents
					//    so it shows on screen, even when you're not looking for his specs

				}, function(err) {
					console.log(err);
				})
			},
		},

		data: {
			agents: [],
			agentinfo: {
				username: '',
				platform: '',
				activity: '',
				microphone: '',
				storylevel: '',
				dzlevel: '',
				gearlevel: '',
				description: '',
				firearms: '',
				stamina: '',
				electronics: ''
			},
			searchparams: {
				platform: 'any',
				activity: 'any',
				microphone: 'any',
				storylevel: 'any',
				dzlevel: 'any',
				gearlevel: 'any',
				search: ''
			},
			shouldShowAgentValidation: false
		},

		computed: {
			validateAgent: function() {
				return {
					username: regexs.username.test(this.agentinfo.username),
					platform: !!this.agentinfo.platform.trim(),
					activity: !!this.agentinfo.activity.trim(),
					microphone: !!this.agentinfo.microphone.trim(),
					storylevel: !!this.agentinfo.storylevel.trim(),
					dzlevel: !!this.agentinfo.dzlevel.trim(),
					gearlevel: !!this.agentinfo.gearlevel.trim(),
					description: this.agentinfo.description.length <= 150,
					firearms: true,
					stamina: true,
					electronics: true
				};
			},
			isAgentFormValid: function() {
				var validation = this.validateAgent;
				return Object.keys(validation).every(function (key) {
					return validation[key];
				});
			},
		},

		ready: function() {
			this.searchAgents();
		},
	});
})();