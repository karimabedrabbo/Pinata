//
//  EmojiCollectionViewCell.swift
//  PINATA
//
//  Created by Karim Abedrabbo on 11/18/18.
//  Copyright © 2018 Karim Abedrabbo. All rights reserved.
//

import UIKit

class EmojiCollectionViewCell: UICollectionViewCell {

    var emojiLabel: UILabel = {
        let emoji = UILabel()
        emoji.translatesAutoresizingMaskIntoConstraints = false
        emoji.font = UIFont.systemFont(ofSize: 60)
        return emoji
    }()
    override init(frame: CGRect) {
        super.init(frame: frame)
        self.addSubview(emojiLabel)
        self.emojiLabel.centerXAnchor.constraint(equalTo: self.centerXAnchor).isActive = true
        self.emojiLabel.centerYAnchor.constraint(equalTo: self.centerYAnchor).isActive = true
    }
    
    
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
}
