Vue.config.devtools = true;
Vue.config.delimiters = ['@{', '}'];
Vue.http.headers.common['X-CSRF-Token'] = document.querySelector('meta[name="csrf-token"]').getAttribute('content');

Vue.filter('hundreds', function(value) {
	if (value >= 1000) {
		return (value / 1000).toFixed(0) + "K";
	} else {
		return value;
	}
});

var app = new Vue({
	el: '#app',

	methods: {
		searchAgents: function() {
			this.$http.post('/app/users').then(function(response) {
				this.$set('agents', response.data.agents);
			},
			function(err) {
				console.log(err);
			});
		}
	},

	data: {
		agents: [],
	},

	ready: function() {
		this.searchAgents();
	},
});