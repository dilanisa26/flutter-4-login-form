class User {
  const User({
    required this.id,
    required this.name,
    required this.balance,
  });

  final int id;
  final String name;
  final int balance;

  factory User.fromJson(Map<String, dynamic> json) {
    final data = json['data'];
    return User(
      id: data['account_id'],
      name: data['name'],
      balance: data['balance'],
    );
  }
}
